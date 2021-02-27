package messages

import (
	"github.com/streadway/amqp"
)

type Handler = func(c Client, d amqp.Delivery) bool

func handlerSkip(c Client, d amqp.Delivery) bool {
	return true
}

type Client struct {
	channel *amqp.Channel
	Queue   string
}

func NewClient(url, exchange, queue string, keys []string) (Client, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return Client{}, err
	}

	ch, err := conn.Channel()

	err = ch.ExchangeDeclare(exchange, amqp.ExchangeTopic, false, true, false, false, nil)
	if err != nil {
		return Client{}, err
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return Client{}, err
	}

	q, err := ch.QueueDeclare(queue, false, true, false, false, nil)
	if err != nil {
		return Client{}, err
	}

	for _, key := range keys {
		err := ch.QueueBind(q.Name, key, exchange, false, nil)
		if err != nil {
			return Client{}, err
		}
	}

	err = ch.QueueBind(q.Name, q.Name, exchange, false, nil)
	if err != nil {
		return Client{}, err
	}

	return Client{
		channel: ch,
		Queue:   q.Name,
	}, err
}

func (c Client) Start(name string, handler Handler) error {
	msgs, err := c.channel.Consume(c.Queue, name, false, false, false, false, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-msgs:
			if handler(c, msg) {
				msg.Ack(false)
			} else {
				msg.Nack(false, false)
			}
		}
	}

	return nil
}

func (c Client) Publish(exchange, key string, msg amqp.Publishing) error {
	return c.channel.Publish(exchange, key, false, false, msg)
}

func (c Client) Stop() {
	c.channel.Close()
}
