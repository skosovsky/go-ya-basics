package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("unable to get list of files", err)

		return
	}
	//  проходим по списку
	for _, file := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, file.Name())
		// печатаем имя элемента
		log.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if file.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("unable to get list of files", err)

		return
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if strings.Contains(path, filter) {
			log.Println(filename)
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
