package messages

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

var (
	timeoutReconnect = 1000 * time.Millisecond
)

type MessageP = amqp.Publishing
type MessageD = amqp.Delivery

type Handler = func(m MessageD) bool

type ClientN struct {
	url string

	errorChannel chan *amqp.Error
	connection   *amqp.Connection
	channel      *amqp.Channel
	closed       bool

	consumers []ConsumerInfo
}

type ConsumerInfo struct {
	Name     string
	Exchange string
	Queue    string
	Keys     []string
	Handler  Handler
}

type messageConsumer func(string)

func NewClientN(url string) *ClientN {
	c := new(ClientN)
	c.url = url
	c.consumers = make([]ConsumerInfo, 0)

	c.connect()
	go c.reconnector()

	return c
}

func (c *ClientN) Send(exchange, key string, msg MessageP) {
	err := c.channel.Publish(
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		msg,      // msg
	)
	logError("Sending message to queue failed", err)
}

func (c *ClientN) Consume(ci ConsumerInfo) {
	log.Println("Registering consumer...")
	deliveries, err := c.registerQueueConsumer(ci)
	log.Println("Consumer registered! Processing messages...")
	c.executeMessageConsumer(err, ci, deliveries, false)
}

func (c *ClientN) Close() {
	log.Println("Closing connection")
	c.closed = true
	c.channel.Close()
	c.connection.Close()
}

func (c *ClientN) reconnector() {
	for {
		err := <-c.errorChannel
		if !c.closed {
			logError("Reconnecting after connection closed", err)

			c.connect()
			c.recoverConsumers()
		}
	}
}

func (c *ClientN) connect() {
	for {
		log.Printf("Connecting to rabbitmq on %s\n", c.url)
		conn, err := amqp.Dial(c.url)
		if err == nil {
			c.connection = conn
			c.errorChannel = make(chan *amqp.Error)
			c.connection.NotifyClose(c.errorChannel)

			log.Println("Connection established!")

			c.openChannel()

			return
		}

		logError("Connection to rabbitmq failed. Retrying in 1 sec... ", err)
		time.Sleep(timeoutReconnect)
	}
}

func (c *ClientN) declareQueue(ci *ConsumerInfo) error {
	err := c.channel.ExchangeDeclare(
		ci.Exchange,        // name
		amqp.ExchangeTopic, // king
		false,              // durable
		true,               // auto-delete
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)

	if err != nil {
		return err
	}

	q, err := c.channel.QueueDeclare(
		ci.Queue, // name
		false,    // durable
		true,     // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return err
	}

	ci.Queue = q.Name

	for _, key := range ci.Keys {
		err := c.channel.QueueBind(ci.Queue, key, ci.Exchange, false, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientN) openChannel() {
	channel, err := c.connection.Channel()
	logError("Opening channel failed", err)
	c.channel = channel
	log.Println("set channel")
}

func (c *ClientN) registerQueueConsumer(ci ConsumerInfo) (<-chan amqp.Delivery, error) {
	c.declareQueue(&ci)

	log.Println("registerQueueConsumer:", ci.Queue, ci.Name)
	msgs, err := c.channel.Consume(
		ci.Queue, // queue
		ci.Name,  // messageConsumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	logError("Consuming messages from queue failed", err)
	return msgs, err
}

func (c *ClientN) executeMessageConsumer(err error, ci ConsumerInfo, deliveries <-chan amqp.Delivery, isRecovery bool) {
	if err == nil {
		if !isRecovery {
			c.consumers = append(c.consumers, ci)
		}
		go func() {
			for delivery := range deliveries {
				ci.Handler(delivery)
			}
		}()
	}
}

func (c *ClientN) recoverConsumers() {
	for i := range c.consumers {
		var ci = c.consumers[i]

		log.Println("Recovering consumer...")
		msgs, err := c.registerQueueConsumer(ci)
		log.Println("Consumer recovered! Continuing message processing...")
		c.executeMessageConsumer(err, ci, msgs, true)
	}
}

func logError(message string, err error) {
	if err != nil {
		log.Printf("%s: %s", message, err)
	}
}
