package mathslice

type Slice []Element

type Element int

// Sum — возвращает сумму элементов.
func Sum(slice Slice) Element {
	var res Element
	for _, s := range slice {
		res += s
	}
	return res
}

// Map — применяет функцию op к каждому элементу.
func Map(slice Slice, op func(Element) Element) {
	for i, s := range slice {
		slice[i] = op(s)
	}
}

// Fold — сворачивает слайс.
func Fold(slice Slice, op func(Element, Element) Element, init Element) Element {
	res := op(init, slice[0])
	for i := 1; i < len(slice); i++ {
		res = op(res, slice[i])
	}
	return res
}
