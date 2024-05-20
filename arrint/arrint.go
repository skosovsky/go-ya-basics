package arrint

import (
	"fmt"
	"strings"
)

// ArrInt описывает слайс целых чисел типа int.
type ArrInt []int

// Add складывает поэлементно два массива и возвращает результирующий слайс.
func Add(aNum, bNum ArrInt) ArrInt {
	length := len(aNum)
	if length-len(bNum) > 0 {
		length = len(bNum)
	}
	cNUm := make(ArrInt, length)
	for i := range length {
		cNUm[i] = aNum[i] + bNum[i]
	}

	return cNUm
}

// Метод String преобразует ArrInt в строку и возвращает её.
func (a ArrInt) String() string {
	out := make([]string, len(a))

	for i, v := range a {
		out[i] = fmt.Sprintf(`<%d>`, v)
	}

	return fmt.Sprintf(`[%s]`, strings.Join(out, ` `))
}
