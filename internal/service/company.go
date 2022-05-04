package service

import (
	"errors"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
)

// for repository level
type CompanyStorager interface {
	CreateCompany(newEntity domain.Company, access domain.AccessMatrix) error
	ReadCompanies(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error)
	UpdateCompanies(sampleCompany domain.Company, filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error)

	// обновить компанию в хранилище
	// запросить лист компаний из хранилища
}

type CompanyService struct {
	CompanyStorage CompanyStorager
}

func (s *CompanyService) CreateCompany(newEntity domain.Company, access domain.AccessMatrix) ([]domain.Company, error) {
	if !access.Create {
		return []domain.Company{}, errors.New("access denied for create company")
	}
	err := s.CompanyStorage.CreateCompany(newEntity, access)
	if err != nil {
		return []domain.Company{}, err
	}
	filter := domain.Filter{Name: []string{newEntity.Name}}
	storedCompany, err := s.CompanyStorage.ReadCompanies(filter, access)
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
	read, err := s.CompanyStorage.ReadCompanies(filter, access)
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
	updated, err := s.CompanyStorage.ReadCompanies(filter, access)
	if err != nil {
		logger.L.Error("check filter: ", filter)
		return []domain.Company{}, err
	}
	return updated, nil
}

func (s *CompanyService) DeleteCompany(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {
	logger.L.Info("STUB: DeleteCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

//////////////////////////

// for middleware:?
// you call the following handler
// it returns a slice: the results of processing
// if the operation is mutable (not read),
// throw them into the queue
func (s *CompanyService) SendJSON(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {
	logger.L.Info("STUB: DeleteCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func NewCompanyService(cStorage CompanyStorager) *CompanyService {
	return &CompanyService{
		CompanyStorage: cStorage,
	}
}
