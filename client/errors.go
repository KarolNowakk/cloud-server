package main

import "errors"

var (
	// errInvalidPath        = errors.New("invalid path provided")
	errFileNotFound       = errors.New("file not found")
	errFileAllreadyExists = errors.New("file allready exists")
)
