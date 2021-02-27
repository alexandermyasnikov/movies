module movies

go 1.16

replace gitlab.com/amyasnikov/movies => /mnt/code/go/movies

require (
	github.com/antchfx/htmlquery v1.2.3
	github.com/go-pg/pg/v10 v10.7.7
	github.com/golang/protobuf v1.4.3
	gitlab.com/amyasnikov/movies v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
