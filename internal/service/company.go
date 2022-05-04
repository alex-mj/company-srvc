package service

import (
	"errors"
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
)

type CompanyStorager interface {
	CreateCompany(newEntity domain.Company) error
	ReadCompanies(filter domain.Filter) ([]domain.Company, error)
	UpdateCompanies(sampleCompany domain.Company, filter domain.Filter) ([]domain.Company, error)
	DeleteCompanies(filter domain.Filter) ([]domain.Company, error)
}

type CompanyService struct {
	CompanyStorage CompanyStorager
}

func (s *CompanyService) CreateCompany(newEntity domain.Company, access domain.AccessMatrix) ([]domain.Company, error) {
	if !access.Create {
		return []domain.Company{}, errors.New("access denied for create company")
	}
	err := s.CompanyStorage.CreateCompany(newEntity)
	if err != nil {
		return []domain.Company{}, err
	}
	filter := domain.Filter{Name: []string{newEntity.Name}}
	storedCompany, err := s.CompanyStorage.ReadCompanies(filter)
	if err != nil {
		logger.L.Errorf("company %s was created, but cannot read it from storage", newEntity.Name)
		return []domain.Company{newEntity}, err
	}
	return storedCompany, nil
}

func (s *CompanyService) ReadCompany(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {
	if !access.Read {
		return []domain.Company{}, errors.New("access denied for Read company")
	}
	read, err := s.CompanyStorage.ReadCompanies(filter)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	return read, nil
}

func (s *CompanyService) UpdateCompany(sampleCompany domain.Company, filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {
	if !access.Update {
		return []domain.Company{}, errors.New("access denied for UPDATE company")
	}
	//saving update list
	read, err := s.CompanyStorage.ReadCompanies(filter)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	_, err = s.CompanyStorage.UpdateCompanies(sampleCompany, filter)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	// return saved list
	returnFilter := domain.Filter{}
	for _, v := range read {
		returnFilter.Code = append(returnFilter.Code, fmt.Sprint(v.Code))
	}
	updated, err := s.CompanyStorage.ReadCompanies(returnFilter)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	return updated, nil
}

func (s *CompanyService) DeleteCompany(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {
	if !access.Delete {
		return []domain.Company{}, errors.New("access denied for DELETE company")
	}
	//deletion operation without the filter will clear the entire list of companies
	if filter.IsEmpty() {
		return []domain.Company{}, errors.New("deletion operation without filter is disabled")
	}
	//saving delete list
	read, err := s.CompanyStorage.ReadCompanies(filter)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	_, err = s.CompanyStorage.DeleteCompanies(filter)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	// return saved list
	return read, nil
}

// TODO: implement it (JSON to MQ)
func (s *CompanyService) SendJSON(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {
	logger.L.Info("STUB: DeleteCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func NewCompanyService(cStorage CompanyStorager) *CompanyService {
	return &CompanyService{
		CompanyStorage: cStorage,
	}
}
