package main

import (
	"log"
)

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

type Option func(*Item)

func Option1(option1 string) Option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}
func Option2(option2 int) Option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

func NewItem(opts ...Option) *Item {
	// инициализируем типовыми значениями
	item := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42, //nolint:mnd // example
	}
	// применяем опции в том порядке, в котором они были заявлены
	for _, opt := range opts {
		opt(item)
	}

	return item
}

func main() {
	// с параметрами по умолчанию
	item1 := NewItem()
	// с применением одной опции
	item2 := NewItem(Option2(70)) //nolint:mnd // example
	// или двух
	item3 := NewItem(Option1("unusual"), Option2(99)) //nolint:mnd // example
	// опции можно заявлять в разном порядке
	item4 := NewItem(Option2(88), Option1("rare")) //nolint:mnd // example

	log.Println(item1, item2, item3, item4)
}
