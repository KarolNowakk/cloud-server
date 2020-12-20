package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getJWTManger() *jwtManager {
	return newJWTManager(&Config{
		Key:      []byte("dslfmoaamciimprcmerpovckmropcmrepocma932k04c"),
		Duration: 5 * time.Minute,
	})
}

func getUser() *User {
	return &User{
		ID:       "afasdf43049x4009jx0",
		Username: "tesUser",
		Email:    "email@email.com",
	}
}

func TestVerify(t *testing.T) {
	m := getJWTManger()
	user := getUser()

	tokenString, err := m.generate(user, randomString(12))

	require.NoError(t, err)

	userID, err := m.verify(tokenString)

	require.NoError(t, err)

	fmt.Println(userID)
}
