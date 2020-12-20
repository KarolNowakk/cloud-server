package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Config is struct for passing config
type Config struct {
	Key      []byte
	Duration time.Duration
}

type userClaims struct {
	jwt.StandardClaims
	UserID string
}

type jwtManager struct {
	config *Config
}

func newJWTManager(conf *Config) *jwtManager {
	return &jwtManager{config: conf}
}

func (m *jwtManager) generate(user *User, tokenID string) (string, error) {
	claims := userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.config.Duration).Unix(),
			Id:        tokenID,
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.config.Key)
}

//verify verifies the access token string and return a user id as string if the token is valid
func (m *jwtManager) verify(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return m.config.Key, nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return "", err
	}

	return claims.UserID, nil
}
