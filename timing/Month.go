package timing

import (
	"log"
	"time"
)

func Month() {
	now := time.Now()
	log.Printf("Current time is %v", now.Format("2006-01-02"))

	currentMonth := time.Now().Month()
	log.Printf("Current month is %v", currentMonth)

	log.Printf("UTC Unix time in second is %v", time.Unix(now.Unix(), 0))
	log.Printf("Current month in UTC Unix time is %v", time.Unix(now.Unix(), 0).Month())
}
