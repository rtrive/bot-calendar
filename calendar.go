package main

import (
	"context"
	"net/http"
	"time"

	log "github.com/rtrive/bot-calendar/log"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const primaryCalendar = "primary"

func getTodayEvent(srv *calendar.Service) []*calendar.Event {
	t := time.Now()

	todayMidnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format(time.RFC3339)
	todayEnd := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 99, t.Location()).Format(time.RFC3339)

	todayEvents, _ := srv.Events.List(primaryCalendar).TimeMin(todayMidnight).TimeMax(todayEnd).ShowDeleted(false).SingleEvents(true).Do()
	return todayEvents.Items
}

func initCalendar(ctx context.Context, client *http.Client) *calendar.Service {
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Error(err)
	}
	return srv
}
