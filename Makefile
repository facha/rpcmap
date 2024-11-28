export PATH := $(PATH):$(HOME)/go/bin

all: generate compile

generate:
	protoc --go_out=server --go-grpc_out=server proto/map.proto
	protoc --go_out=client/go --go-grpc_out=client/go proto/map.proto

compile: generate
	go build -o target/server server/*.go
	go build -o target/go-client client/go/*.go

test: generate
	go test ./...

clean:
	rm -rf target client/go/mapservice server/mapservice

install_tools:
	go get google.golang.org/grpc
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
