FROM golang:1.22-alpine AS build

ENV TZ "Asia/Shanghai"

RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache tzdata

ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOPROXY=https://goproxy.cn/,direct  TZ=Asia/Shanghai

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY docker ./

ARG VERSION
ARG BUILDTIME

RUN  go build -o goframe \
     -ldflags="-s -w -X main.Version=${VERSION} -X main.BuildTime=${BUILDTIME}"  \
     main.go && go clean -cache


# todo  私有镜像
FROM alpine
ENV TZ "Asia/Shanghai"
RUN apk update && apk add --no-cache tzdata
WORKDIR /app
COPY --from=build /app/goframe /app/goframe
COPY --from=build /app/local.config.yaml /app/local.config.yaml




