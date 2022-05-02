package main

import (
	"fmt"

	"github.com/alex-mj/company-srvc/internal/api"
	"github.com/alex-mj/company-srvc/internal/repository"
	"github.com/alex-mj/company-srvc/internal/service"
	"github.com/alex-mj/company-srvc/internal/srv"
)

func main() {
	fmt.Println("Starting COMPANY service...")
	authRepoService := repository.NewAuthRepository()
	userService := service.NewUserService(authRepoService)
	companyStorage := repository.NewPostgresStorage()
	companyService := service.NewCompanyService(companyStorage)
	handlers := api.NewHandler(companyService, userService)

	srv := new(srv.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		//if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		//log.E.Error("error occured while running http server: ", err)
	}
}
