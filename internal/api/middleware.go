package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// авторизация, если пусто то возвращаем матрицу только на чтение
// заполненную проверяем пароль и получаем матрицу привилегий через сервис user (замокать интерфейс userGeter)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	countCtx            = "count"
)

func (h *Handlers) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
	}
	userId, err := h.UserService.GetAccessMatrix(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
