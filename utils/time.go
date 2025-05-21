package utils

import (
	"log"
	"time"
)

func GetLocalBkkTime() time.Time {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println("failed to get bkk time", err)
	}

	return time.Now().In(loc)
}
