package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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

func checkEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv(name)
	return env
}

func main() {
	ctx := context.Background()
	client := initOauth()

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	todayEvents := getTodayEvent(srv)
	for _, event := range todayEvents {
		fmt.Println(event.Summary)
	}
}
