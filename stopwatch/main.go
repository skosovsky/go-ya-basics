package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start  time.Time
	splits []time.Duration
}

func main() {
	swatch := Stopwatch{
		start:  time.Time{},
		splits: nil,
	}
	swatch.Start()

	time.Sleep(1 * time.Second)
	swatch.SaveSplit()

	time.Sleep(500 * time.Millisecond) //nolint:mnd // example
	swatch.SaveSplit()

	time.Sleep(300 * time.Millisecond) //nolint:mnd // example
	swatch.SaveSplit()

	fmt.Println(swatch.GetResults()) //nolint:forbidigo // example
}

func (s *Stopwatch) Start() {
	s.start = time.Now()
	s.splits = nil
}

func (s *Stopwatch) SaveSplit() {
	if s.start.IsZero() {
		return
	}

	s.splits = append(s.splits, time.Since(s.start))
}

func (s *Stopwatch) GetResults() []time.Duration {
	return s.splits
}
