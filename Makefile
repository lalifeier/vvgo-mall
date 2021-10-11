GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)


APP_BASE_RELATIVE_PATH=${shell echo ../../../}
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)


API_PROTO_FILES=$(shell cd ${API_PATH} && find . -name *.proto)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "go-kratos/gvc-" $$0 ":0.1.0"}')

API_PATH=${shell echo ${APP_BASE_RELATIVE_PATH}api/${APP_RELATIVE_PATH}}
PROTO_PATH=${shell echo ${APP_BASE_RELATIVE_PATH}third_party}

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/go-kratos/swagger-api
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get github.com/google/wire/cmd/wire
	go get entgo.io/ent/cmd/ent
	go get -u gorm.io/gorm


.PHONY: api
# generate api proto
api:
	cd ${API_PATH} && protoc --proto_path=. \
	       --proto_path=${PROTO_PATH} \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
				 --validate_out=paths=source_relative,lang=go:. \
	       $(API_PROTO_FILES)

.PHONY: errors
# generate errors code
errors:
	cd ${API_PATH} && protoc --proto_path=. \
               --proto_path=${PROTO_PATH} \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)
validate:
# generate validate code
	cd ${API_PATH} && protoc --proto_path=. \
     					 --proto_path=${PROTO_PATH} \
           		 --go_out=paths=source_relative:. \
           		 --validate_out=paths=source_relative,lang=go:. \
							 $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=${PROTO_PATH} \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: test
# test
test:
	go test -v ./... -cover

.PHONY: run
# run
run:
	cd cmd/$(APP_NAME)/ && go run .

.PHONY: wire
# wire
wire:
	cd cmd/$(APP_NAME)/ && wire

.PHONY: ent
# ent
ent:
	cd internal/data/ && ent generate ./ent/schema

.PHONY: docker
# docker
docker:
	cd ../../ && docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) -t $(DOCKER_IMAGE) .

.PHONY: swagger
# generate swagger
swagger:
	cd ${API_PATH} && protoc --proto_path=. \
	        --proto_path=${PROTO_PATH} \
	        --openapiv2_out . \
	        --openapiv2_opt logtostderr=true \
					--openapiv2_opt json_names_for_fields=false \
           $(API_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
