all: fmt depends build

build:
	go build -race -o ./build/messages ./examples/messages/main.go
	go build -race -o ./build/storage ./service/storage.go
	go build -race -o ./build/parser ./service/parser.go

depends:
	go mod tidy

fmt:
	gofmt -w -s -d .

test:
	go clean -testcache
	go test -race -v ./storage
	go test -race -v ./messages
	go test -race -v ./parser

clean:
	rm -f ./build/*
