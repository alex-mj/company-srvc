package main

import (
	"github.com/alex-mj/company-srvc/internal/api"
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/alex-mj/company-srvc/internal/repository"
	"github.com/alex-mj/company-srvc/internal/service"
	"github.com/alex-mj/company-srvc/internal/srv"
	"github.com/spf13/viper"
)

func main() {

	logger.InitSugar()
	defer logger.SyncForExit()

	if err := initConfig(); err != nil {
		logger.L.DPanic("Check the configuration file (configs/config.yaml):", err)
	}

	logger.L.Info("Starting COMPANY service...")
	// user authentication
	authJWT := repository.NewAuthJWT()
	userService := service.NewUserService(authJWT)

	// company
	companyStorage, err := repository.NewPostgresStorage(getDBConfig())
	if err != nil {
		logger.L.DPanic("Check the data base (postgres): ", err)
	}
	companyService := service.NewCompanyService(companyStorage)

	handlers := api.NewHandler(companyService, userService)

	srv := new(srv.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.L.DPanic("API handler does not start: ", err)
	}
}

func initConfig() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getDBConfig() repository.PostgresConfig {
	return repository.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		UserName: viper.GetString("postgres.user"),
		Password: viper.GetString("postgres.password"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
	}

}
