package main

import "github.com/vferko/golang-brno/db"

type Todo struct {
	Id          int
	Description string
	IsDone      bool
}

func GetAll() []Todo {
	db := db.GetDb()
	rows, err := db.Query("SELECT * FROM")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todoList []Todo

	for rows.Next() {
		todo := Todo{} // same as `var todo TODO`
		err := rows.Scan(&todo.Id, &todo.Description, &todo.IsDone)
		if err != nil {
			panic(err)
		}
		todoList = append(todoList, todo)
	}
	return todoList
}

func GetTodo() Todo {
	return Todo{1, "Golang workshop", false}
}

func SaveTodo(todo Todo) {
	db := db.GetDB()
	_, err := db.Exec("INSERT INTO todo(description, is_done) VALUES($1,$2)")
	if err != nil {
		panic(err)
	}
}
