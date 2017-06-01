package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const spotPriceAPI = "https://api.coinbase.com/v2/prices/spot?currency=USD"

type Amount struct {
	Amount   string
	Currency string
}
type PriceResponse struct {
	Data Amount
}

func CurrentPrice() (string, error) {
	resp, err := http.Get(spotPriceAPI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var latest PriceResponse

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(buf, &latest)
	if err != nil {
		return "", err
	}

	return latest.Data.Amount, nil
}
