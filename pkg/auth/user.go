package auth

//User is user structure for authentication service
type User struct {
	ID                   string
	Email                string
	Password             string
	PasswordConfirmation string
}
