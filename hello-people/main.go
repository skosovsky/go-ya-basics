package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	year := date.Year()

	if year >= 1946 && year <= 1964 {
		fmt.Println("Привет, бумер!")
	} else if year >= 1965 && year <= 1980 {
		fmt.Println("Привет, представитель X!")
	} else if year >= 1981 && year <= 1996 {
		fmt.Println("Привет, миллениал!")
	} else if year >= 1997 && year <= 2012 {
		fmt.Println("Привет, зумер!")
	} else if year >= 2013 {
		fmt.Println("Привет, альфа!")
	} else {
		fmt.Println("Привет!")
	}
}
