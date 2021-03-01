package main

import (
	"encoding/json"
	"log"
	"time"

	"gitlab.com/amyasnikov/movies/messages"
	"gitlab.com/amyasnikov/movies/parser"
)

var (
	rabbitmqURL       = "amqp://guest:guest@127.0.0.1:5672"
	timeoutAfterMovie = 10000
	timeoutAfterTask  = 10 * time.Hour
	moviesCount       = 500
)

func main() {
	opts := parser.Options{
		MediaIndexLimit:     200,
		Lang:                "ru",
		TimeoutMilliSeconds: timeoutAfterMovie,
	}
	p := parser.NewParser(opts)

	client := messages.NewClient(rabbitmqURL)

	handler := func(m messages.MessageD) bool {
		return true
	}

	ci := messages.ConsumerInfo{
		Name:     "",
		Exchange: "movies",
		Queue:    "",
		Keys:     []string{},
		Handler:  handler,
	}

	client.Consume(&ci)

	for {
		for movie := range p.Movies(moviesCount) {
			log.Println(movie.Name)

			json, err := json.Marshal(movie)
			if err != nil {
				log.Println(err)
				continue
			}

			msg := messages.MessageP{
				ReplyTo: ci.Queue,
				Body:    json,
			}
			client.Send("movies", "storage.insert", msg)
		}
		time.Sleep(timeoutAfterTask)
	}
}
