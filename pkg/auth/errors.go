package auth

import "errors"

var (
	//ErrBadCredentials is returned when credentials are not matched
	ErrBadCredentials = errors.New("bad credentials")
)

//ErrValidation is returned while request validating goes wrong
type ErrValidation struct {
	msg string
}

func (e ErrValidation) Error() string {
	return e.msg
}

func (e *ErrValidation) push(validationErr string) {
	e.msg += validationErr + ";"
}
