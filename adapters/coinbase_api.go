package adapters

import (
	"bitcoin-rate/models"
	"encoding/json"
	"io"
	"net/http"
)

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

	var rate models.CoinbaseApiResponse
	err = json.Unmarshal(body, &rate)
	if err != nil {
		return 0, err
	}

	return rate.GetAmountAsFloat64()
}
