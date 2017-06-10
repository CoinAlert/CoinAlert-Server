package main

import (
	"errors"
	"github.com/mikeflynn/go-alexa/skillserver"
	"net/http"
)

var (
	ErrUnsupportedMediaType = errors.New("api: Content-Type unsupported")
	ErrMethodNotAllowed     = errors.New("api: Method not allowed")
	ErrMissingId            = errors.New("api: Id field must not be blank")
	ErrBadAlexaRequest      = errors.New("api: Not a valid Alexa request")
)

func GetHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func PostHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AlexaSkillHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if skillserver.IsValidAlexaRequest(w, r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, ErrBadAlexaRequest.Error(), http.StatusBadRequest)
			return
		}
	})
}
