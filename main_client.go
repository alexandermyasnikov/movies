package main

import (
	"log"
	"strconv"
	"time"

	"gitlab.com/amyasnikov/movies/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = ":8080"

func AboutToSayIt(ctx context.Context, m grpc_api.ServiceMoviesClient, movieKey *grpc_api.MovieKey) (*grpc_api.Movie, error) {
	movie, err := m.GetMovie(ctx, movieKey)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := grpc_api.NewServiceMoviesClient(conn)

	request := &grpc_api.MovieKey{
		Id: "tt12345678",
	}

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		request.Id = "tt" + strconv.Itoa(i)
		response, err := AboutToSayIt(context.Background(), client, request)
		if err != nil {
			panic(err)
		}
		log.Println("request:", request.Id)
		log.Println("response:", response.Data)

	}
}
