package domain

import (
	"fmt"
	"strings"
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

// WHERE NAME IN ('Dreamland') AND CODE in ('2','1') AND WEBSITE IN ('www.dreamland.com') AND...
func (f *Filter) ToSQLWithoutCountry() string {

	if f.IsEmpty() {
		return ""
	}
	var sql, name, website, phone, code string
	sql = "WHERE "
	if len(f.Name) > 0 {
		name = " company.NAME in ('" + strings.Join(f.Name, "', '") + "') "
	}
	if len(f.Website) > 0 {
		website = "AND WEBSITE in ('" + strings.Join(f.Website, "', '") + "') "
	}
	if len(f.Phone) > 0 {
		phone = "AND PHONE in ('" + strings.Join(f.Phone, "', '") + "') "
	}
	if len(f.Code) > 0 {
		code = "AND CODE in ('"
		for v := range f.Code {
			code += "'" + fmt.Sprint(v) + "', "
		}
		code += "') "
	}

	return sql + name + website + phone + code
}

// WHERE NAME IN ("GERMANY", "FRANCE")
func (f *Filter) ToSQLOnLyCountry() string {

	if len(f.Country) == 0 {
		return ""
	}
	return "WHERE COUNTRY.NAME in ('" + strings.Join(f.Country, "', '") + "') "
}

// AND COUNTRY_ID IN ('1', '2)
func (f *Filter) ToSQLCountryCode(codes []int) string {

	if len(codes) == 0 {
		return ""
	}
	sql := " AND COUNTRY_ID in ("
	for v := range codes {
		sql += "'" + fmt.Sprint(v) + "', "
	}
	sql += "') "

	return sql
}
