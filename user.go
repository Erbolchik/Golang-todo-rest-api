package todo

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
