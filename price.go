package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Price struct {
	Current string `json:"currentPrice"`
}

func currentPriceHandler(price *Price) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, err := json.Marshal(price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", buf)
		return
	}
}
