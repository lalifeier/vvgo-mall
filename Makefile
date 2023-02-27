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
DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "vvgo-mall/" $$0 ":0.1.0"}')

API_PATH=${shell echo ${APP_BASE_RELATIVE_PATH}api/${APP_RELATIVE_PATH}}
PROTO_PATH=${shell echo ${APP_BASE_RELATIVE_PATH}third_party}

.PHONY: init api wire config ent run build generate test cover vet lint clean docker openapi app

# init env
init:
	@echo "Installing tools from tools/tools.go"
	@cd tools && cat tools.go |grep _|awk -F '"' '{print $$2}' | xargs -tI % go install %

# generate
generate:
	go generate ./...

# generate api proto
api:
	cd ../../../ && \
	buf generate

# wire
wire:
	wire ./cmd/server

# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=${PROTO_PATH} \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

# ent
ent:
	go generate ./internal/data/ent/schema

# run
run:
	go run ./cmd/server -conf ./configs

# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# test
test:
	go test -v ./app... ./pkg/...

# run coverage tests
cover:
	go test -v ./app... ./pkg/... -coverprofile=coverage.out

# run static analysis
vet:
	go vet

# run lint
lint:
	golangci-lint run -v

# run lint proto
lintproto:
	buf lint

# clean build files
clean:
	go clean
	$(if $(IS_WINDOWS), del "coverage.out", rm -f "coverage.out")

# docker
docker:
	cd ../../../ && \
	docker build -t $(DOCKER_IMAGE) . \
							 -f deploy/build/Dockerfile  \
							 --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH)

# generate OpenAPI v3 doc
openapi:
	cd ../../../ && \
	buf generate --path api/account/service/v1 --template api/account/service/v1/buf.openapi.gen.yaml

# generate swagger
# swagger:
# 	cd ${API_PATH} && protoc --proto_path=. \
# 	        --proto_path=${PROTO_PATH} \
# 	        --openapiv2_out . \
# 	        --openapiv2_opt logtostderr=true \
# 					--openapiv2_opt json_names_for_fields=false \
#            $(API_PROTO_FILES)

# build service all
all:
	make generate;
	make test;
	make run;

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
