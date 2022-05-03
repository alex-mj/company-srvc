package repository

import (
	"fmt"
)

const testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

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
	fmt.Println("STUB: SignIn / TODO: check usr/pwd & return new token")
	return "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", nil
}

func (a *AuthJWTSrvc) CheckToken(userName string, token string) (bool, error) {
	fmt.Println("STUB: CheckToken / TODO:  check token in map[int]string")
	return true, nil
}
