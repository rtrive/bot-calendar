package main

import (
	"testing"

	u "github.com/rtrive/bot-calendar/utility"
)

func TestCheckEnvOk(t *testing.T) {
	t.Setenv("TELEGRAM_API_BOT_KEY", "test")

	got := u.CheckEnv("TELEGRAM_API_BOT_KEY")
	want := "test"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
