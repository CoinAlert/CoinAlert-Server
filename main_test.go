package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	server := "http://localhost:8080"
	url := server + "/api/register"

	r, err := http.Get(url)
	if err != nil {
		t.Errorf("Error on /register GET: %#v\n", err)
	}
	if r.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("GET /register should return %d, got %d\n", http.StatusMethodNotAllowed, r.StatusCode)
	}
}

func TestPostRegister(t *testing.T) {
	server := "http://localhost:8080"
	url := server + "/api/register"

	d := Device{
		Id: "farts",
	}

	j, err := json.Marshal(d)
	if err != nil {
		t.Errorf("Error marshaling device: %#v\n", err)
	}

	jr := bytes.NewReader(j)
	r, err := http.Post(url, "application/json", jr)
	if err != nil {
		t.Errorf("Error on /register POST: %#v\n", err)
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("POST /register should return %d, got %d\n", http.StatusOK, r.StatusCode)
	}

}
