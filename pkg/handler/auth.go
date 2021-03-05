package handler

import (
	"net/http"

	todo "github.com/Erbolchik/Golang-todo-rest-api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) { //registration
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
