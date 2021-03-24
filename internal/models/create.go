package models

import (
	"net/http"

	"github.com/ihor-matiushenko/todo-list/configs"
)

// CreateTodo ...
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &Todo{}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo.Title = r.FormValue("title")

	id := 0

	err = configs.DB.QueryRow("INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id", todo.Title, false).Scan(&id)
	if err != nil {
		panic(err)
	}

	todo.ID = id
}
