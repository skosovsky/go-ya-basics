package main

import "fmt"

func main() {
	var weekTemp = []int{1, 2, 3, 4, 5, 6, 7}
	var newSlc []int
	newSlc = append(newSlc, weekTemp[0:2]...)
	newSlc = append(newSlc, weekTemp[3:]...)

	newSlc[1] = 10
	fmt.Println(weekTemp)
	fmt.Println(newSlc)

	a := []int{1, 2, 3, 4}
	b := a[2:3] // b = [3]
	b = append(b, 7, 8)
	fmt.Println(a, len(a), cap(a)) // [1 2 3 4] 4 4
	fmt.Println(b, len(b), cap(b)) // [3 7 8] 3 4

	s := make([]int, 4)
	_ = append(s[:0], 2, 3, 4, 5)
	fmt.Println(s) // [0 0 2 3]

	slc := []int{1, 2, 3, 4, 5, 6, 7}
	slcd := slc[0:1]                        // 1
	fmt.Println(slcd, len(slcd), cap(slcd)) // 1 7
	slce := slc[0:1:7]                      // 1
	fmt.Println(slce, len(slce), cap(slce)) // 1 3
}
