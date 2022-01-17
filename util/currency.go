package util

// Constants for all supported currencies
const (
	USD = "USD"
	INR = "INR"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR:
		return true
	}
	return false
}
