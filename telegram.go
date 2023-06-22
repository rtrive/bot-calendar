package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron"
	log "github.com/rtrive/bot-calendar/log"
	u "github.com/rtrive/bot-calendar/utility"
	"google.golang.org/api/calendar/v3"
	tele "gopkg.in/telebot.v3"
)

func generateHTML(events []*calendar.Event) string {
	var builder strings.Builder

	builder.WriteString("<ul>")
	builder.WriteString("<li>Start Date</li>")
	builder.WriteString("<li>End Date</li>")
	builder.WriteString("<li>Summary</li>")
	for _, ev := range events {
		builder.WriteString("<li>")
		builder.WriteString(u.GetShortTime(ev.Start.DateTime))
		builder.WriteString("</li>")
		builder.WriteString("<li>")
		builder.WriteString(u.GetShortTime(ev.End.DateTime))
		builder.WriteString("</li>")
		builder.WriteString("<li>")
		builder.WriteString(ev.Summary)
		builder.WriteString("</li>")
	}
	builder.WriteString("</ul>")

	return builder.String()
}

func initBot() {
	log.Debug("Check telegram api key")
	pref := tele.Settings{
		Token:  u.CheckEnv("TELEGRAM_BOT_API_KEY"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	log.Info("Start cron utility")
	cr := cron.New()
	cr.Start()
	log.Info("Cron utility started")
	ctx := context.Background()

	log.Info("Start Google Calendar oauth authentication")
	client := initOauth()
	log.Info("Google Calendar authenticaticated")

	log.Info("Start calendar utility")
	srv := initCalendar(ctx, client)

	log.Info("Init Telegram bot")
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Error(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		log.Debug(fmt.Sprintf("Command start initiated by %s", c.Chat().Username))
		log.Info("Creating calendar and test cron every 1 minute")

		cr.AddFunc("1 * * * * ", func() {
			todayEvent := getTodayEvent(srv)
			var event string
			event += fmt.Sprintf("<pre>")
			event += fmt.Sprintf("| Start Date | End Date | Summary |\n")
			event += fmt.Sprintf("-----------------------------------\n")
			for _, ev := range todayEvent {
				event += fmt.Sprintf("| %s | %s | %s\n", u.GetShortTime(ev.Start.DateTime), u.GetShortTime(ev.End.DateTime), ev.Summary)
			}
			event += fmt.Sprintf("</pre>")
			b.Send(c.Sender(), event, &tele.SendOptions{ParseMode: tele.ModeHTML})
		})

		return c.Send("Bot Started, you will receive some messages, I hope")

	})

	b.Start()
}
