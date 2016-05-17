package main

type Todo struct {
	Id          int
	Description string
	IsDone      bool
}

func GetTodo() Todo {
	return Todo{1, "Golang workshop", false}
}
