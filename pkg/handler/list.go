package handler

import "github.com/gin-gonic/gin"

func (h *Handler) createList(c *gin.Context) {
	id,ok:=c.Get(userCtx)
	if !ok{
		newErrorResponse(c,http.StatusInternalServerError,"user id not found")
		return
	}
	var input todo.TodoList
	if(err):c.BindJson(&input);err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.error())
		return
	}
}
func (h *Handler) getAllLists(c *gin.Context) {
}
func (h *Handler) getListByID(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
