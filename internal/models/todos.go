package models

import (
	"fmt"
	"net/http"

	"github.com/ihor-matiushenko/todo-list/configs"
	"github.com/julienschmidt/httprouter"
)

// Todo ...
type Todo struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Done  bool   `json:"done" db:"done"`
}

// GetTodos ...
func GetTodos() ([]Todo, error) {
	rows, err := configs.DB.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			fmt.Println(err)
			continue
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// AddTodo ...
func AddTodo(w http.ResponseWriter, r *http.Request) {
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

// UpdateTodo ...
func UpdateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()

	title := r.FormValue("title")
	id := ps.ByName("id")

	result, err := configs.DB.Exec("UPDATE todos SET title = $1 WHERE id = $2", title, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}

// CloseTodo ...
func CloseTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := configs.DB.Exec("UPDATE todos SET done = $1 WHERE id = $2", true, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}

// ReopenTodo ...
func ReopenTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := configs.DB.Exec("UPDATE todos SET done = $1 WHERE id = $2", false, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}

// DeleteTodo ...
func DeleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := configs.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
