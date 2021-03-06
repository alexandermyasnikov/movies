package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"gitlab.com/amyasnikov/movies/messages"
	"gitlab.com/amyasnikov/movies/parser"
)

type config struct {
	messagesURL         string
	language            string
	timeoutMsAfterMovie int
	timeoutMsAfterTask  int
	mediaIndexLimit     int
	moviesCount         int
}

func (c *config) init() {
	if c.messagesURL = os.Getenv("MOVIES_PARSER_MESSAGESURL"); c.messagesURL == "" {
		c.messagesURL = "amqp://guest:guest@127.0.0.1:5672"
	}

	if c.language = os.Getenv("MOVIES_PARSER_LANGUAGE"); c.language == "" {
		c.language = "en"
	}

	if val := os.Getenv("MOVIES_PARSER_TIMEOUTMSAFTERMOVIE"); val == "" {
		c.timeoutMsAfterMovie = 10 * 1000
	} else {
		c.timeoutMsAfterMovie, _ = strconv.Atoi(val)
	}

	if val := os.Getenv("MOVIES_PARSER_TIMEOUTMSAFTERTASK"); val == "" {
		c.timeoutMsAfterTask = 1 * 60 * 60 * 1000
	} else {
		c.timeoutMsAfterTask, _ = strconv.Atoi(val)
	}

	if val := os.Getenv("MOVIES_PARSER_MOVIESCOUNT"); val == "" {
		c.moviesCount = 500
	} else {
		c.moviesCount, _ = strconv.Atoi(val)
	}
}

func main() {
	var cfg config
	cfg.init()

	opts := parser.Options{
		MediaIndexLimit:     cfg.mediaIndexLimit,
		Lang:                cfg.language,
		TimeoutMilliSeconds: cfg.timeoutMsAfterMovie,
	}
	p := parser.NewParser(opts)

	client := messages.NewClient(cfg.messagesURL)

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
		for movie := range p.Movies(cfg.moviesCount) {
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
		time.Sleep(time.Duration(cfg.timeoutMsAfterTask) * time.Millisecond)
	}
}
