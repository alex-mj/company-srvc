package api

import (
	"net/http"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) createCompany(c *gin.Context) {
	logger.L.Debug("!!! createCompany !!!")
	access, err := getAccessMatrix(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logger.L.Debug("access %+v", access)
	var newCompany domain.Company
	if err := c.BindJSON(&newCompany); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()+" /check input json")
		return
	}
	logger.L.Debug("newCompany %+v", newCompany)
	if newCompany.Name == "" {
		newErrorResponse(c, http.StatusInternalServerError, "missing company name")
		return
	}
	created, err := h.CompanyService.CreateCompany(newCompany, access)
	logger.L.Debug("h.CompanyService.CreateCompany: ", created, err)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"company": created,
	})
}

func (h *Handlers) readCompanies(c *gin.Context) {
	// filter := domain.Filter{}
	// filter.Name = append(filter.Name, input.Name)
	// filter.Country = append(filter.Country, input.Country)
	// filter.Website = append(filter.Website, input.Website)
	// filter.Phone = append(filter.Phone, input.Phone)
}

func (h *Handlers) readCompany(c *gin.Context) {

}

func (h *Handlers) updateCompany(c *gin.Context) {

}

func (h *Handlers) deleteCompany(c *gin.Context) {

}
