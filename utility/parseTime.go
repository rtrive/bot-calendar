package utility

import (
	"log"
	"time"
)

func GetShortTime(date string) string {

	timeDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	return timeDate.Format(time.Kitchen)

}
