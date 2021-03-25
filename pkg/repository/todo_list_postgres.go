package repository

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres (db *sqlx.DB) *TodoListPostgres{
	return &TodoListPostgres{db:db}
}

func (r *TodoListPostgres) Create(userId int,list todo.TodoList) (int,error){
	tx,err := r.db.Begin()
	if err!=nil{
		return 0,err
	}
	var id int
	createListQuery :=fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id",todoListsTable)
	row:=tx.QueryRow(createListQuery,list.Title,list.Description)
}