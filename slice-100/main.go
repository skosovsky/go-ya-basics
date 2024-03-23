package main

import "fmt"

func main() {
	slc := make([]int, 0, 100)

	for i := 1; i <= 100; i++ {
		slc = append(slc, i)
	}

	slc = append(slc[0:10], slc[90:]...)
	n := len(slc)

	for i := 0; i < n/2; i++ {
		slc[i], slc[n-i-1] = slc[n-i-1], slc[i]
	}

	fmt.Println(slc)
}
