package main

import (
	"testing"

	u "github.com/rtrive/bot-calendar/utility"
)

func TestPartseTimeOk(t *testing.T) {

	got := u.GetShortTime("2023-06-28T10:00:00+02:00")
	//2006-01-02T15:04:05Z07:00
	want := "10:00AM"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
