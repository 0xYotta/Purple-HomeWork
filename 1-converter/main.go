package main

import "fmt"

const USDEUR float64 = 0.92  //06.08.2024
const USDRUB float64 = 85.55 //06.08.2024

func main() {
	EURRUB := USDRUB / USDEUR

	eurAmount := 100.0
	rubAmount := eurAmount * EURRUB

	fmt.Printf("EUR: %.2f = RUB: %.2f", eurAmount, rubAmount)
}
