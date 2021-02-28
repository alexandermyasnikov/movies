all: fmt depends build

build:
	go build -o ./build/messages ./examples/messages/main.go

depends:
	go mod tidy

fmt:
	gofmt -w -s -d .

test:
	go clean -testcache
	go test -race -v ./storage
	go test -race -v ./parser
	go test -race -v ./messages

clean:
	rm -f ./build/*
