package model


type Todo struct{
	Title string
	Done bool
	Priority int
	CreatedAt string
}



var Todos = []Todo{
	{Title: "Buy groceries", Done: false, Priority: 2, CreatedAt: "2024-06-01"},
	{Title: "Read a book", Done: false, Priority: 3, CreatedAt: "2024-06-02"},
	{Title: "Go for a run", Done: false, Priority: 1, CreatedAt: "2024-06-03"},
	{Title: "Call mom", Done: false, Priority: 2, CreatedAt: "2024-06-04"},
	{Title: "Finish project", Done: false, Priority: 1, CreatedAt: "2024-06-05"},
}

func AddTodo(todo Todo) {
	Todos = append(Todos, todo)
}

func GetTodos() []Todo {
	return Todos
}