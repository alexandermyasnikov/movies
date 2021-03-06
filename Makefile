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
	docker build --target movies_storage -t movies_storage .

parser_docker:
	docker build --target movies_parser -t movies_parser .

bot_docker:
	docker build --target movies_bot -t movies_bot .



dev_messages:
	docker run -it --rm -p 5672:5672 -p 15672:15672 rabbitmq:3-management

dev_storage:
	docker run -it --rm -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=dev postgres

dev_storage_test:
	docker run -it --rm -p 5433:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=test postgres
