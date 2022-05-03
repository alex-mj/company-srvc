package service

import (
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
)

// for repository level
type CompanyStorager interface {
	// сохранить компанию в хранилище
	// обновить компанию в хранилище
	// запросить лист компаний из хранилища
}

type CompanyService struct {
	CompanyStorage *CompanyStorager
}

func (s *CompanyService) CreateCompany(filter string) ([]domain.Company, error) {
	fmt.Println("STUB: CreateCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func (s *CompanyService) ReadCompany(filter string) ([]domain.Company, error) {
	fmt.Println("STUB: ReadCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func (s *CompanyService) UpdateCompany(filter string) ([]domain.Company, error) {
	fmt.Println("STUB: UpdateCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func (s *CompanyService) DeleteCompany(filter string) ([]domain.Company, error) {
	fmt.Println("STUB: DeleteCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

// for middleware:
// you call the following handler
// it returns a slice: the results of processing
// if the operation is mutable (not read),
// throw them into the queue
func (s *CompanyService) SendJSON(filter string) ([]domain.Company, error) {
	fmt.Println("STUB: DeleteCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func NewCompanyService(cStorage CompanyStorager) *CompanyService {
	return &CompanyService{
		CompanyStorage: &cStorage,
	}
}
