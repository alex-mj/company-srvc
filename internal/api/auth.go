package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Username string `json: "username" binding: "required"`
	Password string `json: "password" binding: "required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token := "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	// token, err := h.services.Authorization.GenerationToken(input.Username, input.Password)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
