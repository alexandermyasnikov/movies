package main

import (
	"log"
	"strconv"
	"time"

	mg "gitlab.com/amyasnikov/movies/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	host = ":8080"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := mg.NewMoviesStorageClient(conn)

	request := &mg.MovieKey{
		Id: "tt12345678",
	}

	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		request.Id = "tt" + strconv.Itoa(i)
		response, err := client.GetMovie(context.Background(), request)
		if err != nil {
			panic(err)
		}
		log.Println("request:", request.Id)
		log.Println("response:", response.Name)
	}
}
