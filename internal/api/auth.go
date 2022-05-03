package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	UserName string `json: "username" binding: "required"`
	Password string `json: "password" binding: "required"`
}

func (h *Handlers) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		input.UserName = c.Query("username")
		input.Password = c.Query("password")
		if input.UserName == "" || input.Password == "" {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
	token, err := h.UserService.GetToken(input.UserName, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
