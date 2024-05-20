package main

import "log"

func main() {
	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	for i, temp := range &weekTemp { // weekTemp передается по указателю, а не копируется сюда
		log.Println(i, temp)
	}
}
