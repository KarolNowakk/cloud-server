package auth

//Token is token structure for authentication service
type Token struct {
	UserID         string
	ExpirationTime int64
	TokenID        string
	TokenString    string
}
