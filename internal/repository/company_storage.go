package repository

import (
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
)

const (
	companyTable = "company"
	countryTable = "country"
)

func (p *PostgresDB) CreateCompany(company domain.Company) error {
	countryCode, err := p.GetCountryCode(company.Country)
	if err != nil {
		return err
	}
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, country, website, phone) VALUES ($name, $country, $website, $phone)", companyTable)
	row := p.db.QueryRow(query, company.Name, countryCode, company.Website, company.Phone)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (p *PostgresDB) GetCountryCode(country string) (int, error) {
	query := fmt.Sprintf("SELECT code FROM %s WHERE name = $name", countryTable)
	var code int
	fmt.Println("query : ", query)
	err := p.db.Get(&code, query, country)
	if err != nil || code == 0 {
		fmt.Println(">>>>>>>>>: ", err)
		query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($name)", countryTable)
		row := p.db.QueryRow(query, country)
		if err := row.Scan(&code); err != nil {
			return code, err
		}
	}
	return code, nil
}
