package auth

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//Service provides fiel uploading operations
type Service interface {
	Register(ctx context.Context, email, password string) error
	Validate(ctx context.Context, email, password, passwordConfirmation string) error
	Login(ctx context.Context, email, password string) (Token, error)
	Verify(ctx context.Context, tokenString string) (string, error)
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	ValueExists(ctx context.Context, field, value string) bool
	FindUser(ctx context.Context, field, value string) (User, error)
	CreateToken(ctx context.Context, token Token) error
	FindUserByHex(ctx context.Context, id string) (User, error)
	FindUsersTokens(ctx context.Context, id string) ([]Token, error)
}

//NewService returns new upload handler instance
func NewService(r Repository, key []byte, tokenDuration time.Duration) Service {
	return &service{r: r, jwtManager: newJWTManager(&Config{
		Key:      key,
		Duration: tokenDuration,
	})}
}

//Handler handles file upload
type service struct {
	r          Repository
	jwtManager *jwtManager
}

func (s *service) Register(ctx context.Context, email, password string) error {
	newAuthUser := User{}
	newAuthUser.Email = email
	newAuthUser.Password = password

	if err := s.r.CreateUser(ctx, newAuthUser); err != nil {
		return err
	}

	return nil
}

func (s *service) Validate(ctx context.Context, email, password, passwordConfirmation string) error {
	err := ErrValidation{msg: ""}

	if !emailRegex.MatchString(email) {
		err.push("provided email is not valid")
	}
	if len(password) < 10 {
		err.push("password must contain at least 2 uppercase letters, 1 special character, two digits, three lowercase letters and must have length of 10")
	}
	if password != passwordConfirmation {
		err.push("passwords don't match")
	}

	//if values has invalid format ther is no sense to go to db and check if username and email exists
	if len(err.Error()) > 1 {
		return err
	}

	if s.r.ValueExists(ctx, "email", email) {
		err.push("email allready exists")
	}

	if len(err.Error()) > 1 {
		return err
	}

	return nil
}

func (s *service) Login(ctx context.Context, email, password string) (Token, error) {
	user, err := s.r.FindUser(ctx, "email", email)
	if err != nil {
		log.Println(err)
		return Token{}, ErrBadCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println(err)
		return Token{}, ErrBadCredentials
	}

	return s.IssueToken(ctx, user)
}

func (s *service) IssueToken(ctx context.Context, user User) (Token, error) {
	tokenID := randomString(16)

	tokenString, err := s.jwtManager.generate(&user, tokenID)
	if err != nil {
		log.Println(err)
		return Token{}, err
	}

	token := Token{
		TokenID:        tokenID,
		TokenString:    tokenString,
		UserID:         user.ID,
		ExpirationTime: time.Now().Add(s.jwtManager.config.Duration).Unix(),
	}

	if err := s.r.CreateToken(ctx, token); err != nil {
		log.Println(err)
		return Token{}, err
	}

	return token, nil
}

func (s *service) Verify(ctx context.Context, tokenString string) (string, error) {
	id, err := s.jwtManager.verify(tokenString)
	if err != nil {
		return "", err
	}

	_, err = s.r.FindUserByHex(ctx, id)
	if err != nil {
		return "", err
	}

	tokens, err := s.r.FindUsersTokens(ctx, id)
	if err != nil {
		return "", err
	}

	for _, token := range tokens {
		if token.UserID == id {
			return id, nil
		}
	}

	return "", errors.New("unauthenticated")
}
