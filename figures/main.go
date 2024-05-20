package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // Квадрат
	circle                  // Круг
	triangle                // Равносторонний треугольник
)

func area(figure figures) (func(float64) float64, bool) {
	switch figure {
	case square:
		return func(side float64) float64 {
			return side * side
		}, true
	case circle:
		return func(radius float64) float64 {
			return math.Pi * math.Pow(radius, 2) //nolint:mnd // example
		}, true
	case triangle:
		return func(side float64) float64 {
			return math.Pow(side, 2) * math.Sqrt(3) / 4 //nolint:mnd // example
		}, true
	default:
		return nil, false
	}
}

func main() {
	var myFigure figures = 3

	newArea, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка") //nolint:forbidigo // example

		return
	}

	myArea := newArea(10) //nolint:mnd // example
	fmt.Println(myArea)   //nolint:forbidigo // example
}
