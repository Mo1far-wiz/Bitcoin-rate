package models

import "strconv"

type CoinbaseApiResponse struct {
	Data struct {
		Amount   string `json:"amount"`
		Base     string `json:"base"`
		Currency string `json:"currency"`
	} `json:"data"`
}

func (cb *CoinbaseApiResponse) GetAmountAsFloat64() (float64, error) {
	return strconv.ParseFloat(cb.Data.Amount, 64)
}
