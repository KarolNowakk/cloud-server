package storage

import "errors"

var (
	//ErrModelNotFound represents model not found error
	ErrModelNotFound = errors.New("model not found")

	//ErrInvalidLastInsertedID is retruned if last inserted id is not promitive.ObjectId
	ErrInvalidLastInsertedID = errors.New("last inserted is invalid")
)
