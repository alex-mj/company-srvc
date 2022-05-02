package api

import (
	"github.com/alex-mj/company-srvc/domain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CompanyService *domain.CompanyService
	UserService    *domain.UserService
}

func NewHandler(company domain.CompanyService, user domain.UserService) *Handler {
	return &Handler{CompanyService: &company, UserService: &user}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api/v1")
	{
		api.POST("/sign-in", h.signIn)
		companies := api.Group("/companies")
		{
			companies.POST("/", h.createCompany)
			companies.GET("/", h.readCompanies)
			companies.GET("/:id", h.readCompany)
			companies.PUT("/:id", h.updateCompany)
			companies.DELETE("/:id", h.deleteCompany)
		}
	}
	return router
}
