package main

import (
	"encoding/json"
	"log"
	"time"

	mg "gitlab.com/amyasnikov/movies/grpc"
	"gitlab.com/amyasnikov/movies/parser"
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

	p := parser.Parser{}

	client := mg.NewMoviesStorageClient(conn)

	for {
		select {
		case <-time.After(1 * time.Hour):
		case <-time.After(10 * time.Second):
			stats, err := client.GetStats(context.Background(), &mg.Void{})
			if err != nil {
				log.Println(err)
				continue
			}
			if stats.MoviesCount > 50 {
				continue
			}
		}

		for movie := range p.Movies(500) {
			log.Println("movie:", movie.Id, movie.Name, len(movie.Photos))
			json, _ := json.Marshal(movie)
			client.UpdateMovie(context.Background(), &mg.Movie{Id: movie.Id, Json: string(json)})
		}
	}
}
