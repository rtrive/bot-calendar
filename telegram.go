package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func initBot() {
	pref := tele.Settings{
		Token:  checkEnv("TELEGRAM_BOT_API_KEY"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
