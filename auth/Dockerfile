FROM golang:1.17 AS builder

ARG VERSION

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn VERSION=$VERSION make -e build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

VOLUME /data/conf

CMD ["./auth", "-conf", "/data/conf/auth.yaml"]
