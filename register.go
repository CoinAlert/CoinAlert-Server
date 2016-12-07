package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// iOS Specific for now.
type Device struct {
	Id         string
	APNToken   string
	SysVersion string
	SysName    string
	Name       string
	Model      string
}

// Register's an application with the server and/or updates the necessary
// records related to the user and device like push token.
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Require JSON content type.
	content := r.Header.Get("Content-Type")
	if content != "application/json" {
		http.Error(w, ErrUnsupportedMediaType.Error(), http.StatusUnsupportedMediaType)
		return
	}

	// Read the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a temp User to unmarshal
	var d Device
	err = json.Unmarshal(body, &d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = SaveDevice(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Make sure we return 200 OK
	w.WriteHeader(http.StatusOK)
	return
}
