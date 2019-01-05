package util

import (
	"github.com/disiqueira/gocurrency"
	"github.com/shopspring/decimal"
)

func Convert(from, to string, amount float64) float64 {
	f := gocurrency.NewCurrency(from)
	t := gocurrency.NewCurrency(to)
	amt := decimal.NewFromFloat(amount)
	conv, _ := gocurrency.ConvertCurrency(f, t, amt)
	flt, _ := conv.Float64()
	return flt
}