type service

type TodoListService struct{
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService{
	return &TodoListService{repo:repo}
}

func (s *TodoListService) Create(userId int,list todo.TodoList) (int,error)