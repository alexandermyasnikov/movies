package main

import (
	"encoding/json"
	"log"
	"time"

	"gitlab.com/amyasnikov/movies/common"
	"gitlab.com/amyasnikov/movies/messages"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	rabbitmqURL  = "amqp://guest:guest@127.0.0.1:5672"
	token        = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	timeoutMs    = 100
	optionsCount = 7
	similarCount = 2
)

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
	bot, err := tb.NewBot(tb.Settings{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout: time.Duration(timeoutMs) * time.Millisecond},
	})

	if err != nil {
		panic(err)
	}

	client := messages.NewClient(rabbitmqURL)

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
			OptionsCount: optionsCount,
			SimilarCount: similarCount,
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
