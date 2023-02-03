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

## 构建

请在`app/{服务名}/service`下执行：

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
