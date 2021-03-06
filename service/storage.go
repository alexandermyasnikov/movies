package main

import (
	"encoding/json"
	"log"
	"os"

	"gitlab.com/amyasnikov/movies/common"
	"gitlab.com/amyasnikov/movies/messages"
	"gitlab.com/amyasnikov/movies/storage"
)

type config struct {
	messagesURL string
	databaseURL string
}

func (c config) init() {
	if c.messagesURL = os.Getenv("MOVIES_STORAGE_MESSAGESURL"); c.messagesURL == "" {
		c.messagesURL = "amqp://guest:guest@127.0.0.1:5672"
	}

	if c.databaseURL = os.Getenv("MOVIES_STORAGE_DATABASEURL"); c.databaseURL == "" {
		c.databaseURL = "postgresql://postgres:postgres@127.0.0.1/dev?sslmode=disable"
	}
}

func main() {
	var cfg config
	cfg.init()

	db, err := storage.NewDB(cfg.databaseURL)
	if err != nil {
		log.Println(err)
		return
	}

	err = db.CreateSchema()
	if err != nil {
		log.Println(err)
		return
	}

	client := messages.NewClient(cfg.messagesURL)

	handlerInsert := func(m messages.MessageD) bool {
		var req common.APIStorageInsertReq
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

	handlerQuiz := func(m messages.MessageD) bool {
		var req common.APIStorageQuizReq
		err := json.Unmarshal(m.Body, &req)
		if err != nil {
			log.Println(err)
			return false
		}

		quiz, err := db.Quiz(req.OptionsCount, req.SimilarCount)
		if err != nil {
			log.Println(err)
			return false
		}

		json, err := json.Marshal(quiz)
		if err != nil {
			log.Println(err)
			return false
		}

		msg := messages.MessageP{
			ReplyTo:       ciInsert.Queue,
			Body:          json,
			CorrelationId: m.CorrelationId,
		}
		client.Send("movies", m.ReplyTo, msg)

		return true
	}

	ciQuiz := messages.ConsumerInfo{
		Name:     "",
		Exchange: "movies",
		Queue:    "",
		Keys:     []string{"storage.quiz"},
		Handler:  handlerQuiz,
	}

	client.Consume(&ciQuiz)

	select {}
}
