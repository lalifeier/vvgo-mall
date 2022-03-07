FROM golang:alpine

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

ENV TZ Asia/Shanghai
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# WORKDIR /go/src/github.com/vvgo-mall

# ADD go.mod .
# ADD go.sum .
# RUN go mod download
# COPY . .

# RUN go get github.com/golang/protobuf && go get github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/cosmtrek/air && go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest

ENTRYPOINT ["air"]
