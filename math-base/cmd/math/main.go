package main

import (
	"fmt"

	"math-base/pkg/mathslice"
)

func main() {
	slc := mathslice.Slice{1, 2, 3}
	fmt.Println(slc)
	fmt.Println("Сумма слайса: ", mathslice.Sum(slc))

	mathslice.Map(slc, func(i mathslice.Element) mathslice.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", slc)

	fmt.Println("Сумма слайса: ", mathslice.Sum(slc))

	fmt.Println("Свёртка слайса умножением ",
		mathslice.Fold(slc,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		mathslice.Fold(slc,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x + y
			},
			0))

}
