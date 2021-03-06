package handler

import (
	"net/http"

	todo "github.com/Erbolchik/Golang-todo-rest-api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) { //registration
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
