package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func CheckEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv(name)
	return env
}

func main() {

	telegramBotApiKey := CheckEnv("TELEGRAM_BOT_API_KEY")

	pref := tele.Settings{
		Token:  telegramBotApiKey,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
