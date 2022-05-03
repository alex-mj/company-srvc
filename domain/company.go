package domain

//go:generate mockgen -source=company.go -destination=mock/company_mock.go -package=domain

type Company struct {
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

// или заменить на []Company ?
type Filter struct {
	Name    []string `json:"name"`
	Code    []string `json:"code"`
	Country []string `json:"country"`
	Website []string `json:"website"`
	Phone   []string `json:"phone"`
}

type CompanyCreater interface {
	CreateCompany(filter string) ([]Company, error)
}

type CompanyReader interface {
	ReadCompany(filter string) ([]Company, error)
}

type CompanyUpdater interface {
	UpdateCompany(filter string) ([]Company, error)
}

type CompanyDeleter interface {
	DeleteCompany(filter string) ([]Company, error)
}

type QueueMessenger interface {
	SendJSON(filter string) ([]Company, error)
}

type CompanyHandler interface {
	CompanyCreater
	CompanyReader
	CompanyUpdater
	CompanyDeleter
	QueueMessenger
}
