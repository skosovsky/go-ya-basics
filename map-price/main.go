package main

import "fmt"

func main() {
	productsPrice := map[string]int{
		"хлеб":     50,   //nolint:gomnd // it's learning code
		"молоко":   100,  //nolint:gomnd // it's learning code
		"масло":    200,  //nolint:gomnd // it's learning code
		"колбаса":  500,  //nolint:gomnd // it's learning code
		"соль":     20,   //nolint:gomnd // it's learning code
		"огурцы":   200,  //nolint:gomnd // it's learning code
		"сыр":      600,  //nolint:gomnd // it's learning code
		"ветчина":  700,  //nolint:gomnd // it's learning code
		"буженина": 900,  //nolint:gomnd // it's learning code
		"помидоры": 250,  //nolint:gomnd // it's learning code
		"рыба":     300,  //nolint:gomnd // it's learning code
		"хамон":    1500, //nolint:gomnd // it's learning code
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	findPricesOver(productsPrice, 500) //nolint:gomnd // it's learning code
	calculateCostOrder(productsPrice, order)
}

func findPricesOver(productsPrice map[string]int, priceOver int) {
	fmt.Println("Перечень деликатесов:") //nolint:forbidigo // it's learning code
	for product, price := range productsPrice {
		if price >= priceOver {
			fmt.Println("-", product) //nolint:forbidigo // it's learning code
		}
	}
	fmt.Println()
}

func calculateCostOrder(productsPrice map[string]int, productsList []string) {
	var sum int
	for _, product := range productsList {
		sum += productsPrice[product]
	}
	fmt.Println("Стоимость заказа", sum) //nolint:forbidigo // it's learning code
}
