GOFLAGS :=

fmt:
	gofmt -w -s -d .

depends:
	go mod tidy

clean:
	rm -f ./build/*



storage:
	go build $(GOFLAGS) -o ./build/storage ./service/storage.go

parser:
	go build $(GOFLAGS) -o ./build/parser ./service/parser.go

bot:
	go build $(GOFLAGS) -o ./build/bot ./service/bot.go



clean_test:
	go clean -testcache

storage_test:
	go test $(GOFLAGS) -v ./storage

parser_test:
	go test $(GOFLAGS) -v ./parser

messages_test:
	go test $(GOFLAGS) -v ./messages

messages_example:
	go build $(GOFLAGS) -o ./build/messages ./examples/messages/main.go



storage_docker:
	docker build --target storage_parser -t movies_storage .

parser_docker:
	docker build --target movies_parser -t movies_parser .

bot_docker:
	docker build --target movies_bot -t movies_bot .
