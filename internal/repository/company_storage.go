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

func (p *PostgresDB) CreateCompany(company domain.Company) error {
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

	sqlRead := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", countryTable)
	var code int
	err := p.db.Get(&code, sqlRead, country)
	if err != nil {
		logger.L.Warn("sql:", sqlRead, ", error:", err)
	}
	if code == 0 {
		sql := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1)", countryTable)
		row := p.db.QueryRow(sql, country)
		if err := row.Err(); err != nil {
			return code, err
		}
		err = p.db.Get(&code, sqlRead, country)
		if err != nil {
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

func (p *PostgresDB) toSQLWhere(filter domain.Filter) string {

	var sqlCountry, sql string

	// simple field
	sqlFilter := filter.ToSQLWithoutCountry()
	if len(sqlFilter) > 0 {
		sql += sqlFilter
	}
	// country
	if len(filter.Country) > 0 {
		codes, err := p.GetCountryCodes(filter)
		if err != nil {
			logger.L.Warn("sql:", sql, " / error:", err)
			return ""
		}
		sqlCountry = filter.ToSQLCountryCode(codes)
	}
	var and string
	if len(sql) > 0 {
		and = " AND "
	} else {
		and = " WHERE "
	}
	if len(sqlCountry) > 0 {
		sql += and + sqlCountry
	}

	return sql
}

func (p *PostgresDB) ReadCompanies(filter domain.Filter) ([]domain.Company, error) {

	sql := `SELECT distinct
				company.name as name,
				company.code as code,
				country.name as country,
				company.website,
				company.phone
			FROM  company
			INNER JOIN country ON company.country_id = country.id `

	sql += p.toSQLWhere(filter)

	companies := []domain.Company{}
	err := p.db.Select(&companies, sql)
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
	logger.L.Debug("sql:", sql, " / error:", err)
	if err != nil {
		return []int{}, err
	}
	return codes, nil
}

func (p *PostgresDB) toSQLUpdateCompany(sampleCompany domain.Company) string {
	var sql, comma string
	if len(sampleCompany.Name) > 0 {
		sql += fmt.Sprintf("name = '%s'", sampleCompany.Name)
		comma = ", "
	}
	if len(sampleCompany.Country) > 0 {
		countryID, err := p.GetCountryCode(sampleCompany.Country)
		if err != nil {
			logger.L.Warn("sql:", sql, " / error:", err)
			return ""
		}
		sql += comma + fmt.Sprintf(" country_id = %d", countryID)
		comma = ", "
	}
	if len(sampleCompany.Website) > 0 {
		sql += comma + fmt.Sprintf(" website = '%s'", sampleCompany.Website)
		comma = ", "
	}
	if len(sampleCompany.Phone) > 0 {
		sql += comma + fmt.Sprintf(" phone = '%s'", sampleCompany.Phone)
		comma = ", "
	}
	return sql
}

func (p *PostgresDB) UpdateCompanies(sampleCompany domain.Company, filter domain.Filter) ([]domain.Company, error) {
	//UPDATE films SET kind = 'Dramatic' WHERE kind = 'Drama'
	sql := fmt.Sprintf("UPDATE %s SET ", companyTable)
	sql += p.toSQLUpdateCompany(sampleCompany)
	sql += p.toSQLWhere(filter)

	companies := []domain.Company{}
	err := p.db.Select(&companies, sql)
	logger.L.Debug("sql:", sql, " / error:", err)
	if err != nil {
		logger.L.Warn("sql:", sql, " / error:", err)
		return []domain.Company{}, errors.New(err.Error() + " / " + sql)
	}

	return companies, nil
}

func (p *PostgresDB) DeleteCompanies(filter domain.Filter) ([]domain.Company, error) {

	//DELETE FROM films WHERE kind = 'Drama'
	sql := fmt.Sprintf("DELETE FROM %s ", companyTable)
	sql += p.toSQLWhere(filter)

	companies := []domain.Company{}
	err := p.db.Select(&companies, sql)
	logger.L.Debug("sql:", sql, " / error:", err)
	if err != nil {
		logger.L.Warn("sql:", sql, " / error:", err)
		return []domain.Company{}, errors.New(err.Error() + " / " + sql)
	}

	return companies, nil
}
