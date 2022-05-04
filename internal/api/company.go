package api

import (
	"net/http"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) createCompany(c *gin.Context) {

	access, err := getAccessMatrix(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var newCompany domain.Company
	if err := c.BindJSON(&newCompany); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()+" /check input json")
		return
	}

	if newCompany.Name == "" {
		newErrorResponse(c, http.StatusInternalServerError, "missing company name")
		return
	}
	created, err := h.CompanyService.CreateCompany(newCompany, access)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"company": created,
	})
}

func (h *Handlers) readCompanies(c *gin.Context) {

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
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"company": read,
	})

}

func (h *Handlers) updateCompany(c *gin.Context) {

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

	deleted, err := h.CompanyService.DeleteCompany(filter, access)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": deleted,
	})

}
