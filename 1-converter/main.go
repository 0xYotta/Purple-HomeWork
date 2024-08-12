package main

import (
	"fmt"
	"strconv"
	"strings"
)

const usdToEur = 0.92  //08.08.2024
const usdToRub = 85.55 //08.08.2024
const eurToUsd = 1 / usdToEur
const rubToUsd = 1 / usdToRub
const eurToRub = usdToRub / usdToEur
const rubToEur = 1 / eurToRub

type rateMap = map[string]float64     //Aliace to the outer map type
type currencyMap = map[string]rateMap //Aliace to the inner map type

func main() {

	currencyMap := make(currencyMap, 3)
	currencyMap["usd"] = rateMap{
		"eur": usdToEur,
		"rub": usdToRub,
	}

	currencyMap["eur"] = rateMap{
		"usd": eurToUsd,
		"rub": eurToRub,
	}

	currencyMap["rub"] = rateMap{
		"usd": rubToUsd,
		"eur": rubToEur,
	}

	fromCurrency := getInputCurrency("Enter the currency you want to exchange (usd, eur, rub):  ")
	amount := getInputAmount("Enter the amount to exchange:  ")
	toCurrency := getInputCurrency(fmt.Sprintf("Enter the currency you want to exchange into (%s):", getAvailableCurrencies(fromCurrency)), fromCurrency)

	result, rate := calculateExchange(fromCurrency, toCurrency, amount, &currencyMap)

	str := fmt.Sprintf(`
%.2f %s -> %.2f %s
Exchange rate: %.2f 
`,
		amount,
		fromCurrency,
		result,
		toCurrency,
		rate,
	)

	fmt.Println(str)
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

func calculateExchange(fromCurrency, toCurrency string, amount float64, rates *currencyMap) (float64, float64) {
	res := amount * (*rates)[fromCurrency][toCurrency]
	return res, (*rates)[fromCurrency][toCurrency]

}
