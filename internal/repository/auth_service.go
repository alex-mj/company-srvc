package repository

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const (
	testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	tokenTTL  = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
}

// AuthJWTSrvc is a stub for an external service
// We don't do user registration here,
// we don't access the user database,
// we only emulate a fixed list of users.
type AuthJWTSrvc struct {
	users  map[string]string
	tokens map[string]string
}

func NewAuthJWT() *AuthJWTSrvc {
	users := make(map[string]string)
	users["test"] = "test"
	tokens := make(map[string]string)
	tokens["test"] = testToken
	return &AuthJWTSrvc{users: users, tokens: tokens}
}

func (a *AuthJWTSrvc) CreateToken(userName, password string) (string, error) {

	if len(password) == 0 {
		return "", errors.New("password is empty")
	}
	if a.users[userName] != password {
		return "", errors.New("user/password not valid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(viper.GetString("signing.key")))
}

func (a *AuthJWTSrvc) CheckToken(accessToken string) (bool, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(viper.GetString("signing.key")), nil
	})
	if err != nil {
		return false, err
	}
	_, ok := token.Claims.(*tokenClaims)
	if !ok {
		return false, errors.New("token claims are not of type *tokenClaims ")
	}
	return true, nil
}
