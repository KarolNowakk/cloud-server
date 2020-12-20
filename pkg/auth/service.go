package auth

import (
	"cloud/pkg/auth/authpb"
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
	Register(*authpb.RegisterRequest) error
	Validate(*authpb.RegisterRequest) error
	Login(req *authpb.LoginRequest) (*Token, error)
	Verify(tokenString string) (string, error)
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	CreateUser(*User) error
	ValueExists(field, value string) bool
	FindUser(field, value string) (*User, error)
	CreateToken(token *Token) error
	FindUserByHex(id string) (*User, error)
	FindUsersTokens(id string) ([]Token, error)
}

//NewService returns new upload handler instance
func NewService(r Repository, conf *Config) Service {
	return &service{r: r, jwtManager: newJWTManager(conf)}
}

//Handler handles file upload
type service struct {
	r          Repository
	jwtManager *jwtManager
}

func (s *service) Register(req *authpb.RegisterRequest) error {
	newAuthUser := User{}
	newAuthUser.Username = req.Username
	newAuthUser.Email = req.Email
	newAuthUser.Password = req.Password
	newAuthUser.PasswordConfirmation = req.PasswordConfirmation

	if err := s.r.CreateUser(&newAuthUser); err != nil {
		return err
	}

	return nil
}

func (s *service) Validate(req *authpb.RegisterRequest) error {
	err := ErrValidation{msg: ""}

	if len(req.Username) < 8 || len(req.Username) > 30 {
		err.push("username must be between 8 and 30 characters")
	}
	if !emailRegex.MatchString(req.Email) {
		err.push("provided email is not valid")
	}
	if len(req.Password) < 10 {
		err.push("password must contain at least 2 uppercase letters, 1 special character, two digits, three lowercase letters and must have length of 10")
	}
	if req.Password != req.PasswordConfirmation {
		err.push("passwords don't match")
	}

	//if values has invalid format ther is no sense to go to db and check if username and email exists
	if len(err.Error()) > 1 {
		return err
	}

	if s.r.ValueExists("email", req.Email) {
		err.push("email allready exists")
	}
	if s.r.ValueExists("username", req.Username) {
		err.push("username allready exists")
	}

	if len(err.Error()) > 1 {
		return err
	}

	return nil
}

func (s *service) Login(req *authpb.LoginRequest) (*Token, error) {
	user, err := s.r.FindUser("username", req.Username)
	if err != nil {
		log.Println(err)
		return nil, ErrBadCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Println(err)
		return nil, ErrBadCredentials
	}

	return s.IssueToken(user)
}

func (s *service) IssueToken(user *User) (*Token, error) {
	tokenID := randomString(16)

	tokenString, err := s.jwtManager.generate(user, tokenID)
	if err != nil {
		return nil, err
	}

	token := Token{
		TokenID:        tokenID,
		TokenString:    tokenString,
		UserID:         user.ID,
		ExpirationTime: time.Now().Add(s.jwtManager.config.Duration).Unix(),
	}

	if err := s.r.CreateToken(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *service) Verify(tokenString string) (string, error) {
	id, err := s.jwtManager.verify(tokenString)
	if err != nil {
		return "", err
	}

	_, err = s.r.FindUserByHex(id)
	if err != nil {
		return "", err
	}

	tokens, err := s.r.FindUsersTokens(id)
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
