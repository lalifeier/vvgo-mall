//go:build tools
// +build tools

package main

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/bazelbuild/bazelisk"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/go-kratos/kratos/cmd/kratos/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/gnostic"
	_ "github.com/google/gnostic/cmd/protoc-gen-openapi"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/onsi/ginkgo/v2/ginkgo"
	_ "github.com/srikrsna/protoc-gen-gotag"
	_ "github.com/vektra/mockery/v2"
	_ "golang.org/x/tools/cmd/goimports"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
