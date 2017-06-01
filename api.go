package main

import (
	"errors"
	"net/http"
)

var (
	ErrUnsupportedMediaType = errors.New("api: Content-Type unsupported")
	ErrMethodNotAllowed     = errors.New("api: Method not allowed")
	ErrMissingId            = errors.New("api: Id field must not be blank")
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
