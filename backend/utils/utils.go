package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Currencies = map[string]string{
	"USD": "USD",
	"EUR": "EUR",
	"GBP": "GBP",
	"JPY": "JPY",
	"AUD": "AUD",
	"CHF": "CHF",
	"CNY": "CNY",
	"DKK": "DKK",
	"MYR": "MYR",
	"SGD": "SGD",
	"THB": "THB",
}

func IsValidCurrency(currency string) bool {
	if _, ok := Currencies[currency]; ok {
		return true
	}
	return false
}

func GetActiveUser(c *gin.Context) (int64, error) {
	value, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to access resources"})
		return 0, fmt.Errorf("error occured")
	}

	userId, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Encountered an issue"})
		return 0, fmt.Errorf("error occured")
	}

	return userId, nil
}
