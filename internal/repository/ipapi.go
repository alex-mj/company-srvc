package repository

import (
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/fegoa89/ipapi"
)

type IPAPI struct{}

func NewIPAPI() *IPAPI {
	return &IPAPI{}
}

func (i *IPAPI) GetCountry(IP string) (string, error) {
	resp, err := ipapi.FindLocation(IP)
	if err != nil {
		logger.L.Warnf("ipapi.FindLocation: ", err)
		return "", nil
	}
	return resp.Country, nil
}
