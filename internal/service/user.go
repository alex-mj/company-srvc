package service

import (
	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
)

// for repository level
type AuthRepositor interface {
	CheckToken(token string) (bool, error)
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
	logger.L.Info("usr: ", usr, " pwd: ", pwd)
	return s.AuthJWT.CreateToken(usr, pwd)
}

func (s *UserService) GetAccessMatrix(token string) (domain.Access, error) {
	logger.L.Info("STUB: GetAccessMatrix / userID not used -> TODO: 1) 0100 2) map[int]domain.Access")
	access := domain.Access{
		Token: token,
		AccessMatrix: domain.AccessMatrix{
			Create: false,
			Read:   true,
			Update: true,
			Delete: false,
		},
	}
	auth, err := s.AuthJWT.CheckToken(token)
	if err != nil {
		logger.L.Debug(err)
	}
	if auth {
		access.AccessMatrix.Create = true
		access.AccessMatrix.Delete = true
	}
	return access, nil
}

func (s *UserService) ModifyByIP(AccessMatrix domain.AccessMatrix) (domain.AccessMatrix, error) {
	logger.L.Info("STUB: AccessMatrixModifier / TODO:  0**0 => if IP[Ciprus] => 1**1")
	return domain.AccessMatrix{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	}, nil
}
