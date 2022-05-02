package service

import (
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
)

// for repository level
type AuthRepository interface {
	CheckToken(userID int, token string) (bool, error)
	// создать токен
}

type UserService struct {
	AuthRepository *AuthRepository
}

func NewUserService(AuthRepo AuthRepository) *UserService {
	return &UserService{
		AuthRepository: &AuthRepo,
	}
}

func (s *UserService) GetAccessMatrix(userID int) (domain.Access, error) {
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
