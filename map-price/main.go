package main

import (
	"fmt"
	"log"
)

func main() {
	productsPrice := map[string]int{
		"хлеб":     50,   //nolint:mnd // example
		"молоко":   100,  //nolint:mnd // example
		"масло":    200,  //nolint:mnd // example
		"колбаса":  500,  //nolint:mnd // example
		"соль":     20,   //nolint:mnd // example
		"огурцы":   200,  //nolint:mnd // example
		"сыр":      600,  //nolint:mnd // example
		"ветчина":  700,  //nolint:mnd // example
		"буженина": 900,  //nolint:mnd // example
		"помидоры": 250,  //nolint:mnd // example
		"рыба":     300,  //nolint:mnd // example
		"хамон":    1500, //nolint:mnd // example
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	findPricesOver(productsPrice, 500) //nolint:mnd // example
	calculateCostOrder(productsPrice, order)
}

func findPricesOver(productsPrice map[string]int, priceOver int) {
	log.Println("Перечень деликатесов:")
	for product, price := range productsPrice {
		if price >= priceOver {
			fmt.Println("-", product) //nolint:forbidigo // example
		}
	}
	log.Println()
}

func calculateCostOrder(productsPrice map[string]int, productsList []string) {
	var sum int
	for _, product := range productsList {
		sum += productsPrice[product]
	}
	log.Println("Стоимость заказа", sum)
}
