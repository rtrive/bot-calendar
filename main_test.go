package main

import (
	"testing"
)

func TestCheckEnvOk(t *testing.T) {
	t.Setenv("TELEGRAM_API_BOT_KEY", "test")

	got := checkEnv("TELEGRAM_API_BOT_KEY")
	want := "test"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
