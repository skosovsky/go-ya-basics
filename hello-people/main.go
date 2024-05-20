package main //nolint:cyclop // example

import (
	"log"
	"time"
)

func main() {
	date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	year := date.Year()

	switch {
	case year >= 1946 && year <= 1964:
		log.Println("Привет, бумер!")
	case year >= 1965 && year <= 1980:
		log.Println("Привет, представитель X!")
	case year >= 1981 && year <= 1996:
		log.Println("Привет, миллениал!")
	case year >= 1997 && year <= 2012:
		log.Println("Привет, зумер!")
	case year >= 2013: //nolint:mnd // example
		log.Println("Привет, альфа!")
	default:
		log.Println("Привет!")
	}
}
