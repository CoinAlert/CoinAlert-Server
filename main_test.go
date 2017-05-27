package main

import (
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
