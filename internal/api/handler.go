package api

import "github.com/gin-gonic/gin"

type Handler struct {
	//services *service.Service
}

// func NewHandler(services *service.Service) *Handler {
// 	return &Handler{services: services}
// }

func NewHandler() *Handler {
	return &Handler{}
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
