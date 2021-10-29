package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/currency"
)

// Codes returns the supported currency codes for the API
func GetSupportedCodes() {
	currencyService := currency.NewCurrencyService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := currencyService.Codes(&currency.CodesRequest{})
	fmt.Println(rsp, err)
}
