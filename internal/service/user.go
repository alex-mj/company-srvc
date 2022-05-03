package service

import (
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
)

// for repository level
type AuthRepositor interface {
	CheckToken(userName string, token string) (bool, error)
	CreateToken(userName, password string) (string, error)
}

type UserService struct {
	AuthJWT AuthRepositor
}

func NewUserService(AuthJWT AuthRepositor) *UserService {
	return &UserService{
		AuthJWT: AuthJWT,
	}
}

func (s *UserService) GetToken(usr, pwd string) (string, error) {
	return s.AuthJWT.CreateToken(usr, pwd)
}

func (s *UserService) GetAccessMatrix(useNmae string) (domain.Access, error) {
	fmt.Println("STUB: GetAccessMatrix / userID not used -> TODO: 1) 0100 2) map[int]domain.Access")
	return domain.Access{
		Token: "",
		AccessMatrix: domain.AccessMatrix{
			Create: true,
			Read:   true,
			Update: true,
			Delete: true,
		},
	}, nil
}

func (s *UserService) ModifyByIP(AccessMatrix domain.AccessMatrix) (domain.AccessMatrix, error) {
	fmt.Println("STUB: AccessMatrixModifier / TODO:  0**0 => if IP[Ciprus] => 1**1")
	return domain.AccessMatrix{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	}, nil
}
