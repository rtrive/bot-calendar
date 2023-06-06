package main

import "testing"

func TestCheckEnv(t *testing.T) {
	t.Setenv("TELEGRAM_API_BOT_KEY", "test")

	got := CheckEnv("TELEGRAM_API_BOT_KEY")
	want := "test"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
