package api

import (
	"github.com/alex-mj/company-srvc/internal/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.L.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
