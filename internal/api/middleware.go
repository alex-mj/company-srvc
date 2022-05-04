package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/alex-mj/company-srvc/domain"
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	accessCtx           = "accessMatrix"
	countCtx            = "count"
)

func (h *Handlers) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	token := ""
	headerParts := strings.Split(header, " ")
	if len(headerParts) == 2 {
		token = headerParts[1]
	}
	IP := c.ClientIP()
	logger.L.Debugf("userIdentity: %s, %s", IP, token)
	access, err := h.UserService.GetAccessMatrix(token, IP)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	logger.L.Debugf("userIdentity: %+v", access)
	c.Set(accessCtx, access)
}

func getAccessMatrix(c *gin.Context) (domain.AccessMatrix, error) {
	accs, ok := c.Get(accessCtx)
	if !ok {
		return domain.AccessMatrix{}, errors.New("accessMatrix not found")
	}

	accsMtrx, ok := accs.(domain.AccessMatrix)
	if !ok {
		return domain.AccessMatrix{}, errors.New("accessMatrix is of invalid type")
	}

	return accsMtrx, nil
}
