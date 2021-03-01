package main

import (
	"encoding/json"
	"log"

	"gitlab.com/amyasnikov/movies/common"
	"gitlab.com/amyasnikov/movies/messages"
	"gitlab.com/amyasnikov/movies/storage"
)

var (
	storageURL  = "postgresql://postgres:postgres@127.0.0.1/dev?sslmode=disable"
	rabbitmqURL = "amqp://guest:guest@127.0.0.1:5672"
)

func main() {
	db, err := storage.NewDB(storageURL)
	if err != nil {
		log.Println(err)
		return
	}

	err = db.CreateSchema()
	if err != nil {
		log.Println(err)
		return
	}

	client := messages.NewClient(rabbitmqURL)

	handlerInsert := func(m messages.MessageD) bool {
		var req common.StorageMovieAPI
		err := json.Unmarshal(m.Body, &req)
		if err != nil {
			log.Println(err)
			return false
		}

		err = db.Insert(&req)
		if err != nil {
			log.Println(err)
			return false
		}

		return true
	}

	ciInsert := messages.ConsumerInfo{
		Name:     "",
		Exchange: "movies",
		Queue:    "",
		Keys:     []string{"storage.insert"},
		Handler:  handlerInsert,
	}

	client.Consume(&ciInsert)

	select {}
}
