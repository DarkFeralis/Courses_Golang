package main

import "fmt"

func main() {

	const USD_TO_EUR = 0.94
	const USD_TO_RUB = 79

	fmt.Printf("EUR в RUB = %.4f\n", 1/USD_TO_EUR*USD_TO_RUB)

}
