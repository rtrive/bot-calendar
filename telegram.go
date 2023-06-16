package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
	tele "gopkg.in/telebot.v3"
)

func initBot() {
	pref := tele.Settings{
		Token:  checkEnv("TELEGRAM_BOT_API_KEY"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	cr := cron.New()
	cr.Start()
	ctx := context.Background()

	client := initOauth()
	srv := initCalendar(ctx, client)

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		cr.AddFunc("1 * * * *", func() {
			todayEvent := getTodayEvent(srv)
			var event string
			for _, ev := range todayEvent {
				event += fmt.Sprintf("Start Date %s, End Date %s, Summary %s\n", ev.Start.DateTime, ev.End.DateTime, ev.Summary)
			}
			fmt.Println(event)
			b.Send(c.Sender(), event)
		})
		cr.AddFunc("1 * * * *", func() { b.Send(c.Sender(), "test") })

		return c.Send("Bot Started, you will receive some messages, I hope")
	})

	b.Start()
}
