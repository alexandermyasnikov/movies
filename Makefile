all: fmt depends build

build:
	go build -o ./build/messages ./examples/messages/main.go

depends:
	go mod tidy

fmt:
	gofmt -w -s -d .

test:
	go clean -testcache
	go test -v ./storage
	go test -v ./parser

clean:
	rm -f ./build/*
