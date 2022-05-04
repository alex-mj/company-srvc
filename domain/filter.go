package domain

import (
	"fmt"
	"strings"

	"github.com/alex-mj/company-srvc/internal/logger"
)

type Filter struct {
	Name    []string `json:"name"`
	Code    []string `json:"code"`
	Country []string `json:"country"`
	Website []string `json:"website"`
	Phone   []string `json:"phone"`
}

func (f *Filter) IsEmpty() bool {
	return len(f.Name) == 0 &&
		len(f.Code) == 0 &&
		len(f.Country) == 0 &&
		len(f.Website) == 0 &&
		len(f.Phone) == 0
}

func (f *Filter) IsEmptyWithoutCountry() bool {
	return len(f.Name) == 0 &&
		len(f.Code) == 0 &&
		len(f.Website) == 0 &&
		len(f.Phone) == 0
}

// WHERE NAME IN ('Dreamland') AND CODE in ('2','1') AND WEBSITE IN ('www.dreamland.com') AND...
func (f *Filter) ToSQLWithoutCountry() string {

	logger.L.Debug("f.ToSQLWithoutCountry: %+v", f)
	if f.IsEmptyWithoutCountry() {
		return ""
	}
	var sql, name, website, phone, code string
	sql = "WHERE "
	and := ""
	if len(f.Name) > 0 {
		name = " company.NAME in ('" + strings.Join(f.Name, "', '") + "') "
		and = " AND "
	}
	if len(f.Website) > 0 {
		website = and + " WEBSITE in ('" + strings.Join(f.Website, "', '") + "') "
		and = " AND "
	}
	if len(f.Phone) > 0 {
		phone = and + " PHONE in ('" + strings.Join(f.Phone, "', '") + "') "
		and = " AND "
	}
	if len(f.Code) > 0 {
		logger.L.Debug("len(f.Code): ", len(f.Code))
		code = and + " CODE in (" + strings.Join(f.Code, ", ") + ") "
		and = " AND "
	}

	return sql + name + website + phone + code
}

// WHERE NAME IN ("GERMANY", "FRANCE")
func (f *Filter) ToSQLOnLyCountry() string {

	if len(f.Country) == 0 {
		return ""
	}
	var sql, comma string
	for _, v := range f.Country {
		sql += comma + "upper('" + v + "')"
		comma = ","
	}
	return "WHERE upper(COUNTRY.NAME) in (" + sql + ") "
}

// COUNTRY_ID IN ('1', '2)
func (f *Filter) ToSQLCountryCode(codes []int) string {

	if len(codes) == 0 {
		return ""
	}
	sql := " COUNTRY_ID in ("
	comma := ""
	for _, v := range codes {
		logger.L.Debug("v: ", v)
		sql += comma + "'" + fmt.Sprint(v) + "'"
		comma = ","
	}
	sql += ") "

	return sql
}
