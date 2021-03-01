package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"gitlab.com/amyasnikov/movies/messages"
)

func main() {
	url := "amqp://guest:guest@127.0.0.1:5672"

	if len(os.Args) < 3 {
		log.Println("invalid args")
		return
	}

	name := os.Args[1]
	timeoutMs, _ := strconv.Atoi(os.Args[2])
	timeout := time.Duration(timeoutMs) * time.Millisecond

	printHandle := func(m messages.MessageD) bool {
		log.Println("ContentType:", m.ContentType)
		log.Println("CorrelationId:", m.CorrelationId)
		log.Println("ReplyTo:", m.ReplyTo)
		log.Println("ConsumerTag:", m.ConsumerTag)
		log.Println("DeliveryTag:", m.DeliveryTag)
		log.Println("Exchange:", m.Exchange)
		log.Println("RoutingKey:", m.RoutingKey)
		log.Println("Body:", string(m.Body))
		log.Println()
		return true
	}

	switch name {
	case "publish":
		client := messages.NewClient(url)

		ci := messages.ConsumerInfo{
			Name:     "c1",
			Exchange: "exchange_test",
			Queue:    "",
			Keys:     []string{"topic_test"},
			Handler:  printHandle,
		}

		client.Consume(&ci)

		for i := 0; i < 5; i++ {
			log.Println("i:", i)
			client.Send("exchange_test", "topic_test", messages.MessageP{Body: []byte("msg_" + strconv.Itoa(i))})
			time.Sleep(timeout)
		}

		client.Close()
	}

	log.Println("end")
}
