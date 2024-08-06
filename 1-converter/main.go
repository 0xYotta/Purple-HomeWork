package main

import (
	"fmt"
	"strconv"
)

const USDEUR float64 = 0.92  //06.08.2024
const USDRUB float64 = 85.55 //06.08.2024

func main() {
	EURRUB := USDRUB / USDEUR

	eurAmount := getUserInput("Enter EUR amount: ")
	rubAmount := eurAmount * EURRUB

	fmt.Println(fmt.Sprintf("EUR: %.2f = RUB: %.2f", eurAmount, rubAmount))
}

func getUserInput(msg string) float64 {
	var input string
	fmt.Print(msg)
	fmt.Scan(&input)
	fmt.Println()
	res, _ := strconv.ParseFloat(input, 64)
	return res
}

func currencyCalc(amount float64, originalCurrency, targetCurrency string) {

}
