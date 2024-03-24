package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err) //nolint:forbidigo // it's learning code
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(filename) //nolint:forbidigo // it's learning code
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err) //nolint:forbidigo // it's learning code
		return
	}
	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if strings.Contains(path, filter) {
			fmt.Println(filename) //nolint:forbidigo // it's learning code
		}
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}

func main() {
	PrintAllFiles(".")

	PrintAllFilesWithFilter(".", "lesson")
	PrintAllFilesWithFilter(".", "Functions")
}
