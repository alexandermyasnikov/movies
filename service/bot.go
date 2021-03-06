package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"gitlab.com/amyasnikov/movies/common"
	"gitlab.com/amyasnikov/movies/messages"
	tb "gopkg.in/tucnak/telebot.v2"
)

type config struct {
	messagesURL  string
	token        string
	timeoutMs    int
	optionsCount int
	similarCount int
}

func (c *config) init() {
	if c.messagesURL = os.Getenv("MOVIES_BOT_MESSAGESURL"); c.messagesURL == "" {
		c.messagesURL = "amqp://guest:guest@127.0.0.1:5672"
	}

	if c.token = os.Getenv("MOVIES_BOT_TOKEN"); c.token == "" {
		panic("MOVIES_BOT_TOKEN is empty")
	}

	if val := os.Getenv("MOVIES_BOT_TIMEOUTMS"); val == "" {
		c.timeoutMs = 100
	} else {
		c.timeoutMs, _ = strconv.Atoi(val)
	}

	if val := os.Getenv("MOVIES_BOT_OPTIONSCOUNT"); val == "" {
		c.optionsCount = 7
	} else {
		c.optionsCount, _ = strconv.Atoi(val)
	}

	if val := os.Getenv("MOVIES_BOT_SIMILARCOUNT"); val == "" {
		c.similarCount = 2
	} else {
		c.similarCount, _ = strconv.Atoi(val)
	}
}

func processHelp(bot *tb.Bot, m *tb.Message) {
	text := "hello, " + m.Sender.FirstName + "\n"
	text += "Try /start for start quiz."
	bot.Send(m.Chat, text)
}

func processUnknownCommand(bot *tb.Bot, m *tb.Message) {
	text := "Sorry. I don't know this command. Try /help."
	bot.Send(m.Chat, text)
}

func main() {
	var cfg config
	cfg.init()

	bot, err := tb.NewBot(tb.Settings{
		Token: cfg.token,
		Poller: &tb.LongPoller{
			Timeout: time.Duration(cfg.timeoutMs) * time.Millisecond},
	})

	if err != nil {
		panic(err)
	}

	client := messages.NewClient(cfg.messagesURL)

	handlerQuiz := func(m messages.MessageD) bool {
		chat, err := bot.ChatByID(m.CorrelationId)
		if err != nil {
			log.Println(err)
			return false
		}

		var quiz common.APIStorageQuizResp
		err = json.Unmarshal(m.Body, &quiz)
		if err != nil {
			log.Println(err)
			return false
		}

		menuQuiz := &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		btnQuizNext := menuQuiz.Text("Next question")

		menuQuiz.Reply(
			menuQuiz.Row(btnQuizNext),
		)

		p := &tb.Photo{File: tb.FromURL(quiz.Photo)}
		_, err = bot.SendAlbum(chat, tb.Album{p}, menuQuiz)
		if err != nil {
			log.Println(err)
			text := "Sorry. Can not send quiz. Try again."
			bot.Send(chat, text)
			return false
		}

		poll := &tb.Poll{
			Type:          tb.PollQuiz,
			Question:      quiz.Question,
			CorrectOption: quiz.CorrectId,
		}

		for _, o := range quiz.Options {
			poll.AddOptions(o)
		}

		bot.Send(chat, poll, menuQuiz)
		return true
	}

	ciQuiz := messages.ConsumerInfo{
		Name:     "",
		Exchange: "movies",
		Queue:    "",
		Keys:     []string{},
		Handler:  handlerQuiz,
	}

	client.Consume(&ciQuiz)

	handlerStart := func(m *tb.Message) {
		req := common.APIStorageQuizReq{
			OptionsCount: cfg.optionsCount,
			SimilarCount: cfg.similarCount,
		}

		json, err := json.Marshal(req)
		if err != nil {
			log.Println(err)
			return
		}

		msg := messages.MessageP{
			Body:          json,
			ReplyTo:       ciQuiz.Queue,
			CorrelationId: m.Chat.Recipient(),
		}
		client.Send("movies", "storage.quiz", msg)
	}

	bot.Handle("/start", handlerStart)
	bot.Handle("Next question", handlerStart)

	bot.Handle("/help", func(m *tb.Message) {
		processHelp(bot, m)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		processUnknownCommand(bot, m)
	})

	bot.Start()
}
