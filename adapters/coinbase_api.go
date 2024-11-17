package adapters

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

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

type CoinbaseApi struct{}

func (cb *CoinbaseApi) GetBTCRate() (float64, error) {
	resp, err := http.Get("https://api.coinbase.com/v2/prices/BTC-" + Currency + "/spot")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rate CoinbaseApiResponse
	err = json.Unmarshal(body, &rate)
	if err != nil {
		return 0, err
	}

	return rate.GetAmountAsFloat64()
}
