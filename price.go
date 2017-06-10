package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	AlexaResponseVersion      = "1.0"
	OutputSpeechTypePlainText = "PlainText"
	OutputSpeechTypeSSML      = "SSML"
)

type Price struct {
	Current string `json:"currentPrice"`
}

type AlexaResponse struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
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

func alexaHandler(price *Price) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		speechText := fmt.Sprintf("The current price of one Bitcoin is $%s", price.Current)
		os := OutputSpeech{Type: OutputSpeechTypePlainText, Text: speechText}
		resp := Response{OutputSpeech: os, ShouldEndSession: true}
		ar := AlexaResponse{Version: AlexaResponseVersion, Response: resp}

		buf, err := json.Marshal(ar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", buf)
		return
	}
}
