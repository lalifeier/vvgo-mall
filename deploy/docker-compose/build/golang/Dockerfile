# build stage
FROM golang:latest as build

ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
ENV CGO_ENABLED=0

WORKDIR /app

COPY . /app
RUN make build

# production stage
FROM alpine as production
# FROM centurylink/ca-certs

WORKDIR /app
COPY --from=build /app .

EXPOSE 8081

ENTRYPOINT ["./server"]
