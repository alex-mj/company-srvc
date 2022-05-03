package service

import (
	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/spf13/viper"
)

const CYPRUS string = "CY"

// for repository level
type AuthRepositor interface {
	CheckToken(token string) (bool, error)
	CreateToken(userName, password string) (string, error)
}

type ContryIdentifier interface {
	GetCountry(IP string) (string, error)
}

type UserService struct {
	AuthJWT AuthRepositor
	IPAPI   ContryIdentifier
}

func NewUserService(AuthJWT AuthRepositor, IPAPI ContryIdentifier) *UserService {
	return &UserService{
		AuthJWT: AuthJWT,
		IPAPI:   IPAPI,
	}
}

func (s *UserService) GetToken(usr, pwd string) (string, error) {
	logger.L.Info("usr: ", usr, " pwd: ", pwd)
	return s.AuthJWT.CreateToken(usr, pwd)
}

func (s *UserService) GetAccessMatrix(token, IP string) (domain.AccessMatrix, error) {
	access := domain.AccessMatrix{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	}
	// option 1 (cyprus):
	accessCyprus := s.GetAccessMatrixCyprus(IP)
	// option 2 (jwt):
	accessJWT := s.GetAccessMatrixJWT(token)

	// assembly:
	access.Create = access.Create && accessCyprus.Create && accessJWT.Create
	access.Read = access.Read && accessCyprus.Read && accessJWT.Read
	access.Update = access.Update && accessCyprus.Update && accessJWT.Update
	access.Delete = access.Delete && accessCyprus.Delete && accessJWT.Delete

	return access, nil
}

func (s *UserService) GetAccessMatrixCyprus(IP string) domain.AccessMatrix {
	accessCyprus := domain.AccessMatrix{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	}
	if viper.GetBool("access.option.cyprus") {
		accessCyprus.Create = false
		accessCyprus.Delete = false
		if len(IP) > 7 {
			country, err := s.IPAPI.GetCountry(IP)
			if err != nil {
				logger.L.Debug(err)
				return accessCyprus
			}
			if country == CYPRUS {
				accessCyprus.Create = true
				accessCyprus.Delete = true
			}
		}
	}
	return accessCyprus
}
func (s *UserService) GetAccessMatrixJWT(token string) domain.AccessMatrix {
	accessJWT := domain.AccessMatrix{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	}
	if viper.GetBool("access.option.jwt") {
		accessJWT.Create = false
		accessJWT.Delete = false
		if len(token) > 0 {
			authOK, err := s.AuthJWT.CheckToken(token)
			if err != nil {
				logger.L.Debug(err)
				return accessJWT
			}
			if authOK {
				accessJWT.Create = true
				accessJWT.Delete = true
			}
		}
	}
	return accessJWT
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
