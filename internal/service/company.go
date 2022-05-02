package service

import (
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
)

// for repository level
type CompanyStorage interface {
	// сохранить компанию в хранилище
	// обновить компанию в хранилище
	// запросить лист компаний из хранилища
}

type CompanyService struct {
	CompanyStorage *CompanyStorage
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

// for middleware
// вызываешь следующий обработчик
// он возвращает массив: результаты обработки
// если операция мутабельная (не чтение),
// то перебрасываем их в очередь
func (s *CompanyService) SendJSON(filter string) ([]domain.Company, error) {
	fmt.Println("STUB: DeleteCompany / filter not used -> TODO: 1) DB 2) filter")
	return []domain.Company{}, nil
}

func NewCompanyService(cStorage CompanyStorage) *CompanyService {
	return &CompanyService{
		CompanyStorage: &cStorage,
	}
}
