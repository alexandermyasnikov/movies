package messages

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

var (
	url          = "amqp://guest:guest@127.0.0.1:5672"
	timeoutSleep = 100 * time.Millisecond
)

func TestSendSelf(t *testing.T) {
	client := NewClient(url)

	var mu sync.Mutex
	msgs := []string{}

	ci := ConsumerInfo{
		Name:     "c1",
		Exchange: "exchange_test",
		Queue:    "",
		Keys:     []string{"#"},
		Handler: func(m MessageD) bool {
			mu.Lock()
			defer mu.Unlock()
			msgs = append(msgs, string(m.Body))
			return true
		},
	}

	client.Consume(&ci)

	time.Sleep(timeoutSleep)

	client.Send("exchange_test", "test", MessageP{Body: []byte("hello1")})
	client.Send("exchange_test", "test", MessageP{Body: []byte("hello2")})
	client.Send("exchange_test", "test", MessageP{Body: []byte("hello3")})

	time.Sleep(timeoutSleep)

	client.Close()

	mu.Lock()
	defer mu.Unlock()

	msgsEx := []string{"hello1", "hello2", "hello3"}

	if !reflect.DeepEqual(msgs, msgsEx) {
		t.Errorf("msgs = %v; want %v", msgs, msgsEx)
	}
}

func TestManyConsumers(t *testing.T) {
	client := NewClient(url)

	var mu sync.Mutex
	msgs1 := []string{}
	msgs2 := []string{}
	msgs3 := []string{}

	client.Consume(&ConsumerInfo{
		Name:     "c1",
		Exchange: "exchange_test",
		Queue:    "",
		Keys:     []string{"topic_test1"},
		Handler: func(m MessageD) bool {
			mu.Lock()
			defer mu.Unlock()
			msgs1 = append(msgs1, string(m.Body))
			return true
		},
	})

	client.Consume(&ConsumerInfo{
		Name:     "c2",
		Exchange: "exchange_test",
		Queue:    "",
		Keys:     []string{"topic_test2"},
		Handler: func(m MessageD) bool {
			mu.Lock()
			defer mu.Unlock()
			msgs2 = append(msgs2, string(m.Body))
			return true
		},
	})

	client.Consume(&ConsumerInfo{
		Name:     "c3",
		Exchange: "exchange_test",
		Queue:    "",
		Keys:     []string{"topic_test1", "topic_test2", "topic_test3"},
		Handler: func(m MessageD) bool {
			mu.Lock()
			defer mu.Unlock()
			msgs3 = append(msgs3, string(m.Body))
			return true
		},
	})

	time.Sleep(timeoutSleep)

	client.Send("exchange_test", "topic_test1", MessageP{Body: []byte("hello1")})
	client.Send("exchange_test", "topic_test2", MessageP{Body: []byte("hello2")})
	client.Send("exchange_test", "topic_test3", MessageP{Body: []byte("hello3")})
	client.Send("exchange_test", "topic_test1", MessageP{Body: []byte("hello4")})
	client.Send("exchange_test", "topic_test2", MessageP{Body: []byte("hello5")})
	client.Send("exchange_test", "topic_test3", MessageP{Body: []byte("hello6")})

	time.Sleep(timeoutSleep)

	client.Close()

	mu.Lock()
	defer mu.Unlock()

	msgs1Ex := []string{"hello1", "hello4"}
	msgs2Ex := []string{"hello2", "hello5"}
	msgs3Ex := []string{"hello1", "hello2", "hello3", "hello4", "hello5", "hello6"}

	if !reflect.DeepEqual(msgs1, msgs1Ex) {
		t.Errorf("msgs1 = %v; want %v", msgs1, msgs1Ex)
	}

	if !reflect.DeepEqual(msgs2, msgs2Ex) {
		t.Errorf("msgs2 = %v; want %v", msgs2, msgs2Ex)
	}

	if !reflect.DeepEqual(msgs3, msgs3Ex) {
		t.Errorf("msgs3 = %v; want %v", msgs3, msgs3Ex)
	}
}
