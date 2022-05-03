package repository

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestWriteCountry(t *testing.T) {

	p := initPostgres()

	want := 1
	actual, _ := p.GetCountryCode("Dreamland")

	require.Equal(t, want, actual, "check id of country: dreamland != 0")
}

func initPostgres() *PostgresDB {
	viper.Set("postgres.host", "localhost")
	viper.Set("postgres.port", "5432")
	viper.Set("postgres.user", "postgres")
	viper.Set("postgres.password", "qwerty5432")
	viper.Set("postgres.dbname", "postgres")
	viper.Set("postgres.sslmode", "disable")
	p, err := NewPostgresStorage(GetDBConfig())
	if err != nil {
		panic(err)
	}
	return p
}
