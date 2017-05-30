package main

import "errors"

var (
	ErrUnsupportedMediaType = errors.New("api: Content-Type unsupported")
	ErrMethodNotAllowed     = errors.New("api: Method not allowed")
	ErrMissingId            = errors.New("api: Id field must not be blank")
)
