all: fmt depends build

build:
	go build -o ./build/messages ./examples/messages/main.go

depends:
	go mod tidy

fmt:
	gofmt -w -s -d .

clean:
	rm -f ./build/*
