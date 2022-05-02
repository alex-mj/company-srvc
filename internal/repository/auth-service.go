package repository

import (
	"fmt"
)

type AuthRepoSrvc struct {
	tokens map[int64]string
}

func NewAuthRepository() *AuthRepoSrvc {
	return &AuthRepoSrvc{make(map[int64]string)}
}

func (a *AuthRepoSrvc) CheckToken(userID int, token string) (bool, error) {
	fmt.Println("STUB: CheckToken / TODO:  check token in map[int]string")
	return true, nil
}
