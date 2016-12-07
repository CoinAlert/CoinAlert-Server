package main

import "errors"

var (
	ErrUnsupportedMediaType = errors.New("api: Content-Type unsupported")
)
