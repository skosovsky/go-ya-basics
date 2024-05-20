package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	log.Println("Interaction counter")

	cnt := 0
	for {
		log.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		log.Printf("User input %d lines\n", cnt)
	}
}

func f(cnt *int) {
	*cnt++
}
