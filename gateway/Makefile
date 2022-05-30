GOPATH:=$(shell go env GOPATH)
CONFIG_PROTO_FILES=$(shell find internal -name conf.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: config
# generate config proto
config:
	protoc --proto_path=./internal \
 	       --go_out=paths=source_relative:./internal \
	       $(CONFIG_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go mod tidy
	go get -d github.com/google/wire/cmd/wire@latest
	go generate ./...
