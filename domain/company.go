package domain

//go:generate mockgen -source=company.go -destination=mock/company_mock.go -package=domain

type Company struct {
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type CompanyCreater interface {
	CreateCompany(newEntity Company, access AccessMatrix) ([]Company, error)
}

type CompanyReader interface {
	ReadCompany(filter Filter, access AccessMatrix) ([]Company, error)
}

type CompanyUpdater interface {
	UpdateCompany(sampleCompany Company, filter Filter, access AccessMatrix) ([]Company, error)
}

type CompanyDeleter interface {
	DeleteCompany(filter Filter, access AccessMatrix) ([]Company, error)
}

type QueueMessenger interface {
	SendJSON(filter Filter, access AccessMatrix) ([]Company, error)
}

type CompanyHandler interface {
	CompanyCreater
	CompanyReader
	CompanyUpdater
	CompanyDeleter
	QueueMessenger
}
