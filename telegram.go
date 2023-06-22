package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/robfig/cron"
	log "github.com/rtrive/bot-calendar/log"
	u "github.com/rtrive/bot-calendar/utility"
	"google.golang.org/api/calendar/v3"
	tele "gopkg.in/telebot.v3"
)

func generateTable(events []*calendar.Event) string {
	tw := table.NewWriter()
	colTitleStartDate := "Start Date"
	colTitleEndDate := "End Date"
	colTitleSummary := "Summary"
	tableHeader := table.Row{colTitleStartDate, colTitleEndDate, colTitleSummary}
	tw.SetTitle("Calendar")
	tw.SetStyle(table.StyleRounded)
	tw.AppendHeader(tableHeader)
	for _, ev := range events {
		tw.AppendRow(table.Row{u.GetShortTime(ev.Start.DateTime), u.GetShortTime(ev.End.DateTime), ev.Summary})
	}
	return fmt.Sprintf("<pre>%s</pre>", tw.Render())
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
			message := generateTable(todayEvent)
			b.Send(c.Sender(), message, &tele.SendOptions{ParseMode: tele.ModeHTML})
		})

		return c.Send("Bot Started, you will receive some messages, I hope")
	})

	b.Start()
}
