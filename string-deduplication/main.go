package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(removeDuplicates(input)) //nolint:forbidigo // it's learning code
}

func removeDuplicates(input []string) []string {
	noDuplicates := make(map[string]struct{}, len(input))
	noDuplicatesSlc := make([]string, 0)

	for _, v := range input {
		if _, ok := noDuplicates[v]; !ok {
			noDuplicatesSlc = append(noDuplicatesSlc, v)
		}
		noDuplicates[v] = struct{}{}
	}

	return noDuplicatesSlc
}
