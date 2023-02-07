GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

APP_BASE_RELATIVE_PATH=${shell echo ../../../}
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)

API_PROTO_FILES=$(shell cd ${API_PATH} && find . -name *.proto)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "go-kratos/gvc-" $$0 ":0.1.0"}')

API_PATH=${shell echo ${APP_BASE_RELATIVE_PATH}api/${APP_RELATIVE_PATH}}
PROTO_PATH=${shell echo ${APP_BASE_RELATIVE_PATH}third_party}

.PHONY: init api wire config ent run build generate test cover vet lint docker swagger app

# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/bazelbuild/bazelisk@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/google/gnostic@latest
	go install github.com/google/wire/cmd/wire@latest
	go install entgo.io/ent/cmd/ent@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# generate api proto
api:
	cd ../../../ && \
	buf generate

# wire
wire:
	cd cmd/$(APP_NAME)/ && wire

# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=${PROTO_PATH} \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

# ent
ent:
	cd internal/data/ && go generate ./ent/schema

# run
run:
	cd cmd/$(APP_NAME)/ && go run .

# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# generate
generate:
	go generate ./...

# test
test:
	go test -v ./... -cover

# run coverage tests
cover:
	go test -v ./... -coverprofile=coverage.out

# run static analysis
vet:
	go vet

# run lint
lint:
	golangci-lint run

# docker
docker:
	cd ../../ && \
	docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(DOCKER_IMAGE) .

# generate swagger
swagger:
	cd ${API_PATH} && protoc --proto_path=. \
	        --proto_path=${PROTO_PATH} \
	        --openapiv2_out . \
	        --openapiv2_opt logtostderr=true \
					--openapiv2_opt json_names_for_fields=false \
           $(API_PROTO_FILES)

# build service app
app: api wire config ent build

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
