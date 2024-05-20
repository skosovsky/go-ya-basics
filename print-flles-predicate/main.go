package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintAllFilesWithFilterClosure(path string, filter string) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err) //nolint:forbidigo // example

			return
		}
		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if strings.Contains(filename, filter) {
				fmt.Println(filename) //nolint:forbidigo // example
			}
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	walk(path)
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err) //nolint:forbidigo // example

			return
		}
		for _, file := range files {
			filename := filepath.Join(path, file.Name())
			fmt.Println(filename) //nolint:forbidigo // example
			if predicate(filename) {
				fmt.Println(filename) //nolint:forbidigo // example
			}

			if file.IsDir() {
				walk(filename)
			}
		}
	}
	walk(path)
}

// containsDot возвращает все пути, содержащие точки.
func containsDot(s string) bool {
	return strings.Contains(s, ".")
}

func main() {
	// PrintAllFilesWithFilterClosure("../", "num")

	PrintFilesWithFuncFilter("../", containsDot)
}
