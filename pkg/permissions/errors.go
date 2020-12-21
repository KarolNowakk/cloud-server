package permissions

import "errors"

var (
	//ErrInvalidPath is returned if path is invalid
	ErrInvalidPath = errors.New("invalid path provided")

	//ErrPermissionDenied is returned if user has no permissions to data
	ErrPermissionDenied = errors.New("permission denied")
)
