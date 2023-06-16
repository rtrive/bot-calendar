package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
	srv := initCalendar(ctx, client)

	todayEvents := getTodayEvent(srv)
	for _, event := range todayEvents {
		fmt.Println(event.Summary)
	}
}
