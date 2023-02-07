# vvgo-mall

- 基于 golang 微服务框架 [go-kratos](https://go-kratos.dev)

## 技术栈

- [Kratos](https://go-kratos.dev)
- [Consul](https://www.consul.io)
- [OpenTelemetry](https://opentelemetry.io)
- [Prometheus](https://prometheus.io)
- [Wire](https://github.com/google/wire)
- [Ent](https://entgo.io)
- [MySQL](https://www.mysql.com)
- [MongoDB](https://www.mongodb.com)
- [Redis](https://redis.io)
- [copier](https://github.com/jinzhu/copier)
- [Casbin](https://casbin.org)
- [OAuth2](https://github.com/golang/oauth2)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [Sarama](https://github.com/Shopify/sarama)
- [AWS SDK](https://aws.github.io/aws-sdk-go-v2/docs)
- [Alibaba Cloud SDK](https://help.aliyun.com/document_detail/122613.html)
- [Qiniu Cloud SDK](https://developer.qiniu.com/kodo/1238/go)
- [tencentyun SDK](https://cloud.tencent.com/document/product/436/31215)

## 前置环境

### Protobuf

```bash
sudo apt install protobuf-compiler
```

## Make 构建

在`app/{服务名}/service`下执行

- 初始化开发环境

  ```bash
  make init
  ```

- 生成 API 的 go 代码

  ```bash
  make api
  ```

- 生成 API 的 OpenAPI v3 文档

  ```bash
  make openapi
  ```

- 生成配置结构代码

  ```bash
  make config
  ```

- 生成 wire 代码

  ```bash
  make wire
  ```

- 生成 ent 代码

  ```bash
  make ent
  ```

- 构建程序

  ```bash
  make build
  ```

- 构建 Docker 镜像

  ```bash
  make docker
  ```

## Bazel 构建

使用[bazel.build](https://bazel.build/)进行服务器程序的构建。

如何安装 bazel.build 的文档，请参考官方文档：<https://bazel.build/install>。

在根目录下执行命令：

- 更新 GO 依赖库引入的构建配置文件 repos.bzl

  ```bash
  bazel run //:gazelle-update-repos
  ```

- 拉取依赖项，生成配置文件 BUILD.bazel

  ```bash
  bazel run //:gazelle
  ```

- 构建单个服务

  ```bash
  bazel build //app/admin/service/cmd/server:server
  # 或者
  bazel build //:admin-service
  ```

- 运行单个服务

  ```bash
  bazel run //app/admin/service/cmd/server:server
  # 或者
  bazel run //:admin-service
  ```

- 单个服务生成本地 Docker 镜像

  ```bash
  bazel run //:admin-service-image
  ```

- 单个服务生成 Docker 镜像 tar 文件

  ```bash
  bazel build //:admin-service-image
  ```

- 单个服务推送 Docker 镜像到 DockerHub

  ```bash
  bazel run //:admin-service-image-push
  ```

- 构建全部服务

  ```bash
  bazel build //...
  ```
