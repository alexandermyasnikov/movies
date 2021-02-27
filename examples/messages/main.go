package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"gitlab.com/amyasnikov/movies/messages"
)

type Request struct {
	Id   string `json:"id"`
	Time string `json:"time"`
	Text string `json:"text"`
}

func main() {
	url := "amqp://guest:guest@127.0.0.1:5672"

	if len(os.Args) < 3 {
		log.Println("invalid args")
		return
	}

	name := os.Args[1]
	timeout, _ := strconv.Atoi(os.Args[2])

	printHandle := func(c messages.Client, d amqp.Delivery) bool {
		log.Println("ContentType:", d.ContentType)
		log.Println("CorrelationId:", d.CorrelationId)
		log.Println("ReplyTo:", d.ReplyTo)
		log.Println("ConsumerTag:", d.ConsumerTag)
		log.Println("DeliveryTag:", d.DeliveryTag)
		log.Println("Exchange:", d.Exchange)
		log.Println("RoutingKey:", d.RoutingKey)
		log.Println("Body:", string(d.Body))
		log.Println()
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		return true
	}

	switch name {
	case "publish":
		client, err := messages.NewClient(url, "exchange_movies", "", []string{})
		if err != nil {
			panic(err)
		}

		defer client.Stop()

		go func() {
			for {
				req := Request{
					Id:   "id_1",
					Time: time.Now().String(),
					Text: fmt.Sprintf("text_%v", time.Now().UTC().UnixNano()),
				}

				json, err := json.Marshal(req)
				if err != nil {
					log.Println(err)
					continue
				}

				log.Println("publish:", string(json))
				msg := amqp.Publishing{
					ContentType:   "application/json",
					Body:          json,
					CorrelationId: strconv.FormatInt(time.Now().UTC().UnixNano(), 16),
					ReplyTo:       client.Queue,
				}

				client.Publish("exchange_movies", "log.print", msg)
				client.Publish("exchange_movies", "worker.reply", msg)
				time.Sleep(time.Duration(timeout) * time.Millisecond)

			}
		}()

		err = client.Start("publisher", printHandle)
		if err != nil {
			panic(err)
		}
	case "log.print":
		client, err := messages.NewClient(url, "exchange_movies", "", []string{"log.print"})
		if err != nil {
			panic(err)
		}

		defer client.Stop()

		err = client.Start("logger", printHandle)
		if err != nil {
			panic(err)
		}
	case "reply":
		client, err := messages.NewClient(url, "exchange_movies", "", []string{"worker.reply"})
		if err != nil {
			panic(err)
		}

		defer client.Stop()

		replyHandle := func(c messages.Client, d amqp.Delivery) bool {
			time.Sleep(time.Duration(timeout) * time.Millisecond)
			if d.RoutingKey == c.Queue {
				return true
			}
			if d.Body == nil {
				log.Println("Error, no message body!")
				return true
			}
			log.Println("reply: publish:", d.ReplyTo, d.RoutingKey, d.CorrelationId)

			msg := amqp.Publishing{
				ContentType:   "application/json",
				Body:          d.Body,
				CorrelationId: "reply_" + d.CorrelationId,
				ReplyTo:       client.Queue,
			}

			c.Publish("exchange_movies", d.ReplyTo, msg)
			return true
		}

		err = client.Start("replier", replyHandle)
		if err != nil {
			panic(err)
		}
	case "worker":
		client, err := messages.NewClient(url, "exchange_movies", "worker", []string{"worker.reply"})
		if err != nil {
			panic(err)
		}

		defer client.Stop()

		replyHandle := func(c messages.Client, d amqp.Delivery) bool {
			time.Sleep(time.Duration(timeout) * time.Millisecond)
			if d.RoutingKey == c.Queue {
				return true
			}
			if d.Body == nil {
				log.Println("Error, no message body!")
				return true
			}
			log.Println("worker: publish:", d.ReplyTo, d.RoutingKey, d.CorrelationId)

			msg := amqp.Publishing{
				ContentType:   "application/json",
				Body:          d.Body,
				CorrelationId: "reply_" + d.CorrelationId,
				ReplyTo:       client.Queue,
			}

			c.Publish("exchange_movies", d.ReplyTo, msg)
			return true
		}

		err = client.Start("worker", replyHandle)
		if err != nil {
			panic(err)
		}
	}

	log.Println("end")
}
