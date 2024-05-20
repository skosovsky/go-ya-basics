package main

import (
	"log"
)

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func NextDay(day Weekday) Weekday {
	return (day % 7) + 1 //nolint:mnd // example
}

func main() {
	var today = Sunday
	tomorrow := NextDay(today)
	log.Println("today =", today, "tomorrow =", tomorrow)
}
