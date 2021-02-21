module movies

go 1.16

replace gitlab.com/amyasnikov/movies => /mnt/code/go/movies

require (
	github.com/golang/protobuf v1.4.3
	github.com/mattn/go-sqlite3 v1.14.6
	gitlab.com/amyasnikov/movies v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
