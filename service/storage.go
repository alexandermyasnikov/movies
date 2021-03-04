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
