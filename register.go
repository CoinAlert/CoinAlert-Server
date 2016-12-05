package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//	"gopkg.in/mgo.v2"
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

}
