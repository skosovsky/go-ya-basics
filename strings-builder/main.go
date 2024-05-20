package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	word := strings.Builder{}

	for range 50 {
		_, err := fmt.Fprintf(&word, "%v", math.NaN())
		if err != nil {
			return
		}
	}
	word.WriteString("... BATMAN!")

	log.Printf("%s\n", &word)
}
