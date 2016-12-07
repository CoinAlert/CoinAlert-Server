package main

import "errors"

var (
	ErrUnsupportedMediaType = errors.New("api: Content-Type unsupported")
	ErrMissingId            = errors.New("api: Id field must not be blank")
)
