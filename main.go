package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())

	if now.Month() == time.January && now.Day() == 1 {
		newYear = time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	}

	duration := newYear.Sub(now)
	days := int(duration.Hours() / 24)

	if days == 0 {
		fmt.Println("С Новым годом!")
	} else {
		fmt.Printf("До Нового года осталось %d дней\n", days)
	}
}
