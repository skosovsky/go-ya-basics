package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 7

	fmt.Println(findPairs(nums, k))
}

func findPairs(nums []int, k int) [][]int {
	idxNum := make(map[int]int)
	var pairs [][]int

	for i, num := range nums {
		if j, ok := idxNum[k-num]; ok {
			pairs = append(pairs, []int{i, j})
		}
		idxNum[num] = i
	}

	return pairs
}
