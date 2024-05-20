package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrPanicRecovered = errors.New("panic recovered")

type input struct {
	a       int
	b       int
	counter int
}

func fn(a, b, counter int) {
	if counter == 0 {
		panic("counter equals 0")
	}

	fn(b, a/b, counter-1)
}

func test(a, b, counter int) (err error) { //nolint:varnamelen // example
	defer func() {
		if reasonPanic := recover(); reasonPanic != nil {
			err = fmt.Errorf("%s: %w", reasonPanic, ErrPanicRecovered)
		}
	}()

	fn(a, b, counter)

	return nil
}

func main() {
	inputs := []input{
		{10, 5, 3},
		{100, 7, 10},
		{1, 1, 1000},
	}

	for _, pars := range inputs {
		log.Printf("(%d, %d, %d) => %v\r\n", pars.a, pars.b, pars.counter, test(pars.a, pars.b, pars.counter))
	}
}
