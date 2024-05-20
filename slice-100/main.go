package main

import (
	"log"
)

func main() {
	slc := make([]int, 0, 100) //nolint:mnd // example

	for i := 1; i <= 100; i++ {
		slc = append(slc, i)
	}

	slc = append(slc[0:10], slc[90:]...)
	n := len(slc)

	for i := range n / 2 {
		slc[i], slc[n-i-1] = slc[n-i-1], slc[i]
	}

	log.Println(slc)
}
