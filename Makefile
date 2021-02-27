all: grpc fmt depends build

build:
	go build -o main_storage ./main_storage.go
	go build -o main_client  ./main_client.go
	go build -o main_parser  ./main_parser.go

depends:
	go mod tidy

grpc:
	protoc -I. --go_out=plugins=grpc:. ./grpc/api.proto

fmt:
	gofmt -w -s -d .
