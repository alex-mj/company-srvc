package api

import (
	"github.com/alex-mj/company-srvc/domain"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	CompanyService domain.CompanyHandler
	UserService    domain.UserHandler
}

func NewHandler(company domain.CompanyHandler, user domain.UserHandler) *Handlers {
	return &Handlers{CompanyService: company, UserService: user}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api/v1")
	{
		api.POST("/sign-in", h.signIn)
		companies := api.Group("/companies", h.userIdentity)
		{
			companies.POST("/", h.createCompany)
			companies.GET("/", h.readCompanies)
			companies.GET("/:id", h.readCompanies)
			companies.PUT("/:id", h.updateCompany)
			companies.DELETE("/:id", h.deleteCompany)
		}
	}
	return router
}
