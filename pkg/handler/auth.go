package handler

import (
	"fmt"
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
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("Test test")
}
