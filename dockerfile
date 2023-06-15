FROM golang:alpine as builder

WORKDIR /app
# 源
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ginsample-api main.go


FROM alpine:3.18
# 设置时区
RUN apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone
RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash

WORKDIR /app
# 复制到工作区
COPY --from=builder /app/ ./
# COPY --from=builder /work/config ./config
# 执行
CMD ["./ginsample-api"]