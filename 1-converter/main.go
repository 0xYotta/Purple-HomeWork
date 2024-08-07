package main

import (
	"fmt"
	"strconv"
	"strings"
)

const usdToEur = 0.92  //08.08.2024
const usdToRub = 85.55 //08.08.2024

func main() {
	fromCurrency := getInputCurrency("Enter the currency you want to exchange (usd, eur, rub):  ")
	amount := getInputAmount("Enter the amount to exchange:  ")
	toCurrency := getInputCurrency(fmt.Sprintf("Enter the currency you want to exchange into (%s):", getAvailableCurrencies(fromCurrency)), fromCurrency)

	result := calculateExchange(fromCurrency, toCurrency, amount)
	fmt.Printf("Amount after exchange: %.2f %s\n", result, toCurrency)
}

func getInputCurrency(prompt string, exclude ...string) string {
	for {
		fmt.Println(prompt)
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToLower(currency)
		if isValidCurrency(currency, exclude...) {
			return currency
		}
		fmt.Println("Invalid currency. Please try again.")
	}
}

func getAvailableCurrencies(exclude string) string {
	available := ""
	if exclude != "usd" {
		available += "usd, "
	}
	if exclude != "eur" {
		available += "eur, "
	}
	if exclude != "rub" {
		available += "rub, "
	}
	return available
}

func isValidCurrency(currency string, exclude ...string) bool {

	if currency != "usd" && currency != "eur" && currency != "rub" {
		return false
	}
	for _, ex := range exclude {
		if currency == ex {
			return false
		}
	}
	return true
}

func getInputAmount(prompt string) float64 {
	for {
		fmt.Println(prompt)
		var amountStr string
		fmt.Scan(&amountStr)
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Invalid amount. Please try again.")
	}
}

func calculateExchange(fromCurrency, toCurrency string, amount float64) float64 {
	var result float64
	switch fromCurrency {
	case "usd":
		if toCurrency == "eur" {
			result = amount * usdToEur
		} else if toCurrency == "rub" {
			result = amount * usdToRub
		}
	case "eur":
		if toCurrency == "usd" {
			result = amount / usdToEur
		} else if toCurrency == "rub" {
			result = amount / usdToEur * usdToRub
		}
	case "rub":
		if toCurrency == "usd" {
			result = amount / usdToRub
		} else if toCurrency == "eur" {
			result = amount / usdToRub * usdToEur
		}
	}
	return result
}
