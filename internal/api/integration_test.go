package api

import (
	"testing"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/repository"
	"github.com/alex-mj/company-srvc/internal/service"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestCRUDCompany(t *testing.T) {

	// init
	companySrvc := initCompanieSrvc()
	accessFull := domain.AccessMatrix{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	}
	companies := []domain.Company{
		{
			Name:    "Company1",
			Code:    0,
			Country: "cntr1",
			Website: "some-url.com",
			Phone:   "000-000-000",
		},
		{
			Name:    "Company2",
			Code:    0,
			Country: "cntr2",
			Website: "some-url.com",
			Phone:   "000-000-000",
		},
		{
			Name:    "Company3",
			Code:    0,
			Country: "cntr2",
			Website: "some-url.com",
			Phone:   "000-000-000",
		},
	}

	filterCntr1 := domain.Filter{
		Country: []string{"cntr1"},
	}
	filterCntr2 := domain.Filter{
		Country: []string{"cntr2"},
	}
	filterName3 := domain.Filter{
		Name: []string{"Company3"},
	}
	filterName1 := domain.Filter{
		Name: []string{"Company1"},
	}
	filterAll := domain.Filter{
		Country: []string{"cntr1", "cntr2"},
	}
	// clear on EXIT
	companySrvc.DeleteCompany(filterCntr1, accessFull)
	companySrvc.DeleteCompany(filterCntr2, accessFull)
	// Test: CreateCompany
	// 1) write 3 companies ( 1 country (cntr1) + 2 countries (cntr2) )
	// read them in a list (by country)
	for _, cmpn := range companies {
		companySrvc.CreateCompany(cmpn, accessFull)
	}

	// check
	{
		readCmpns, _ := companySrvc.ReadCompany(filterCntr1, accessFull)
		want := 1
		actual := len(readCmpns)
		require.Equal(t, want, actual, "Try write 3 company, and read 1 by country (cntr1)")
	}
	{
		readCmpns, _ := companySrvc.ReadCompany(filterCntr2, accessFull)
		want := 2
		actual := len(readCmpns)
		require.Equal(t, want, actual, "Try write 3 company, and read 2 by country (cntr2)")

	}
	// Test: DeleteCompanyByName 
	// 2) remove one of them by name
	// *) read all of them (we see 2 of them)
	companySrvc.DeleteCompany(filterName3, accessFull)
	// check
	{
		readCmpns, _ := companySrvc.ReadCompany(filterAll, accessFull)
		want := 2
		actual := len(readCmpns)
		require.Equal(t, want, actual, "delete company3, and read 2")
	}
	// TestUpdateCompanyByName
	// 3) update one of them (by name)
	// *) read by country (we see them 2)
	sampleCompany := domain.Company{
		Country: "cntr2",
	}
	companySrvc.UpdateCompany(sampleCompany, filterName1, accessFull)
	// check
	{
		readCmpns, _ := companySrvc.ReadCompany(filterCntr2, accessFull)
		want := 2
		actual := len(readCmpns)
		require.Equal(t, want, actual, " 1 delete, 1 update -> read 2 by country (cntr2)")
	}
	// TestDeleteAllCompany
	// 4) delete all companies (by country)
	// *) read all (see 0)
	companySrvc.DeleteCompany(filterCntr1, accessFull)
	companySrvc.DeleteCompany(filterCntr2, accessFull)
	// check
	{
		readCmpns, _ := companySrvc.ReadCompany(filterAll, accessFull)
		want := 0
		actual := len(readCmpns)
		require.Equal(t, want, actual, " all delete")
	}
	// clear on EXIT
	companySrvc.DeleteCompany(filterCntr1, accessFull)
	companySrvc.DeleteCompany(filterCntr2, accessFull)
}

// ---- init's ---
func initUserSrvc() *service.UserService {
	// user authentication
	authJWT := repository.NewAuthJWT()
	ipapi := repository.NewIPAPI()
	return service.NewUserService(authJWT, ipapi)
}

func initCompanieSrvc() *service.CompanyService {
	companyStorage := initPostgres()
	return service.NewCompanyService(companyStorage)
}

func initPostgres() *repository.PostgresDB {
	viper.Set("postgres.host", "localhost")
	viper.Set("postgres.port", "5432")
	viper.Set("postgres.user", "postgres")
	viper.Set("postgres.password", "qwerty5432")
	viper.Set("postgres.dbname", "postgres")
	viper.Set("postgres.sslmode", "disable")
	p, err := repository.NewPostgresStorage(repository.GetDBConfig())
	if err != nil {
		panic(err)
	}
	return p
}
