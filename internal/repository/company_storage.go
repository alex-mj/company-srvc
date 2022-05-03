package repository

import (
	"errors"
	"fmt"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/jmoiron/sqlx"
)

const (
	companyTable = "company"
	countryTable = "country"
)

func (p *PostgresDB) CreateCompany(company domain.Company, access domain.AccessMatrix) error {
	countryCode, err := p.GetCountryCode(company.Country)
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("INSERT INTO %s (name, country_id, website, phone) SELECT $1, $2, $3, $4", companyTable)
	sql += fmt.Sprintf(" WHERE not exists(select 1,2,3,4 from %s WHERE name = $5)", companyTable)
	row := p.db.QueryRow(sql, company.Name, countryCode, company.Website, company.Phone, company.Name)
	if err := row.Err(); err != nil {
		logger.L.Warn("sql:", sql, ", error:", err)
		return errors.New(err.Error() + " / " + sql)
	}
	return nil
}

func (p *PostgresDB) GetCountryCode(country string) (int, error) {

	sql := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", countryTable)
	var code int
	err := p.db.Get(&code, sql, country)

	if err != nil || code == 0 {
		logger.L.Warn("sql:", sql, ", error:", err)
		sql := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", countryTable)
		row := p.db.QueryRow(sql, country)
		if err := row.Err(); err != nil {
			logger.L.Warn("sql:", sql, ", error:", err)
			return code, err
		}
	}

	return code, nil
}

func (p *PostgresDB) GetCompany(companyName string) error {
	sql := fmt.Sprintf("SELECT code FROM %s WHERE name = $name", companyTable)
	var company domain.Company
	rows, err := p.db.Query(sql, company.Name)
	if err != nil {
		logger.L.Error("sql:", sql, "error:", err)
		return errors.New(err.Error() + " / " + sql)
	}
	if err = sqlx.StructScan(rows, &company); err != nil {
		return err
	}
	return nil
}

func (p *PostgresDB) ReadCompanies(filter domain.Filter, access domain.AccessMatrix) ([]domain.Company, error) {

	sqlCountry := ""
	if len(filter.Country) > 0 {
		codes, err := p.GetCountryCodes(filter)
		if err != nil {
			return []domain.Company{}, err
		}
		sqlCountry = filter.ToSQLCountryCode(codes)
	}

	sql := `SELECT distinct
				company.name as name,
				company.code as code,
				country.name as country,
				company.website,
				company.phone
			FROM  company
			INNER JOIN country ON company.country_id = country.id `
	sql += filter.ToSQLWithoutCountry() + sqlCountry
	companies := []domain.Company{}

	err := p.db.Select(&companies, sql)
	logger.L.Debug("sql:", sql, " / error:", err)
	if err != nil {
		logger.L.Warn("sql:", sql, " / error:", err)
		return []domain.Company{}, errors.New(err.Error() + " / " + sql)
	}

	return companies, nil
}

func (p *PostgresDB) GetCountryCodes(filter domain.Filter) ([]int, error) {

	sql := fmt.Sprintf("SELECT id FROM %s ", countryTable) + filter.ToSQLOnLyCountry()
	codes := []int{}
	err := p.db.Select(&codes, sql)
	if err != nil {
		return []int{}, err
	}
	return codes, nil
}
