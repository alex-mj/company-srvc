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
	logger.L.Debug("!!! readCompanies !!!")
	access, err := getAccessMatrix(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filter := domain.Filter{}
	filter.Name = c.QueryArray("name")
	filter.Code = c.QueryArray("code")
	filter.Country = c.QueryArray("country")
	filter.Website = c.QueryArray("website")
	filter.Phone = c.QueryArray("phone")

	id := c.Param("id")
	if id != "" {
		filter.Code = append(filter.Code, id)
	}
	read, err := h.CompanyService.ReadCompany(filter, access)
	logger.L.Debug("h.CompanyService.readCompanies: ", read, err)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"company": read,
	})

}

func (h *Handlers) updateCompany(c *gin.Context) {
	logger.L.Debug("!!! updateCompany !!!")
	access, err := getAccessMatrix(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filter := domain.Filter{}
	filter.Name = c.QueryArray("name")
	filter.Code = c.QueryArray("code")
	filter.Country = c.QueryArray("country")
	filter.Website = c.QueryArray("website")
	filter.Phone = c.QueryArray("phone")

	id := c.Param("id")
	if id != "" {
		filter.Code = append(filter.Code, id)
	}

	var sampleCompany domain.Company
	if err := c.BindJSON(&sampleCompany); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()+" /check input json")
		return
	}
	logger.L.Debug("sampleCompany %+v", sampleCompany)
	updated, err := h.CompanyService.UpdateCompany(sampleCompany, filter, access)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"updated": updated,
	})

}

func (h *Handlers) deleteCompany(c *gin.Context) {
	logger.L.Debug("!!! deleteCompany !!!")
	access, err := getAccessMatrix(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filter := domain.Filter{}
	filter.Name = c.QueryArray("name")
	filter.Code = c.QueryArray("code")
	filter.Country = c.QueryArray("country")
	filter.Website = c.QueryArray("website")
	filter.Phone = c.QueryArray("phone")

	id := c.Param("id")
	if id != "" {
		filter.Code = append(filter.Code, id)
	}
	logger.L.Debug("delete filter %+v", filter)
	deleted, err := h.CompanyService.DeleteCompany(filter, access)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": deleted,
	})

}
