package gocurrency

import (
	"strconv"
	"math"
)

//Convert the specified currency to EUR; Returns (-)infinite float64 if currency doesn't exist,
func ConvertToEUR(currency string, value float64) (float64) {
	return value / GetRate(currency)
}

//Convert EUR to a specific currency; Returns 0 if currency doesn't exist
func ConvertEURTo(targetCurrency string, value float64) (float64) {
	return GetRate(targetCurrency) * value
}

//Convert from a specific currency to a specific currency; Returns (-)infinite or 0 if currency doesn't exist.
func ConvertFromTo(fromCurrency, toCurrency string, value float64) (float64) {
	return ConvertEURTo(toCurrency, ConvertToEUR(fromCurrency, value))
}

//Format a float64 number
func FormatNumber(value float64, precision int) (string) {
	return strconv.FormatFloat(value, 'f', precision, 64)
}

//Check if value is infinite or 0 (float default)
func IsNumberValid(value float64) (bool) {
	return !math.IsInf(value, 0) && value != 0
}

//Return the rate of a currency; Returns default float64 if it can't find the currency
func GetRate(currency string) (float64) {
	RefreshIfRequired()
	rate, _:= ExchangeRates[currency]
	return rate
}

