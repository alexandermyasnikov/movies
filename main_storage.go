package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"gitlab.com/amyasnikov/movies/grpc"
	"google.golang.org/grpc"
)

type ServiceMessage struct {
}

var port = ":8080"

func (ServiceMessage) GetMovie(ctx context.Context, r *grpc_api.MovieKey) (*grpc_api.Movie, error) {
	log.Println("Request id:", r.Id)
	response := &grpc_api.Movie{
		Data: r.Id + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
	}
	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Serving requests...")
	server := grpc.NewServer()
	var ServiceMessage ServiceMessage
	grpc_api.RegisterServiceMoviesServer(server, ServiceMessage)
	server.Serve(listen)
}
