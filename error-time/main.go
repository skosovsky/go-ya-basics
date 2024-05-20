package main

import (
	"errors"
	"fmt"
	"time"
)

type TimeError struct {
	Time time.Time
	Text string
}

func (t TimeError) Error() string {
	return fmt.Sprintf("%v: %v", t.Time.Format(`2006/01/02 15:04:05`), t.Text)
}

func NewTimeError(text string) TimeError {
	return TimeError{
		Time: time.Now(),
		Text: text,
	}
}

func testFunc(num int) error {
	if num == 0 {
		return NewTimeError("num = 0")
	}

	return nil
}

func main() {
	if err := testFunc(0); err != nil {
		var timeError TimeError
		if ok := errors.As(err, &timeError); ok {
			fmt.Println(timeError.Time, timeError.Text) //nolint:forbidigo // example
		} else {
			fmt.Println(err) //nolint:forbidigo // example
		}
	}
}
