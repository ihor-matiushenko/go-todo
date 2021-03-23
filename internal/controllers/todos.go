package controllers

import (
	"fmt"
	"net/http"
	"sort"

	// ...
	_ "github.com/lib/pq"

	"github.com/ihor-matiushenko/todo-list/configs"
	"github.com/ihor-matiushenko/todo-list/internal/models"
	"github.com/julienschmidt/httprouter"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos, err := models.GetTodos()

	if err != nil {
		fmt.Println(err)
		return
	}

	sort.SliceStable(todos, func(i, j int) bool {
		return todos[i].ID < todos[j].ID
	})

	w.WriteHeader(http.StatusOK)
	configs.TPL.ExecuteTemplate(w, "todo-list.html", todos)
}

// CreateTodo ...
func CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	models.AddTodo(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// UpdateTodo ...
func UpdateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	models.UpdateTodo(w, r, ps)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// CloseTodo ...
func CloseTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	models.CloseTodo(w, r, ps)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ReopenTodo ...
func ReopenTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	models.ReopenTodo(w, r, ps)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteTodo ...
func DeleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	models.DeleteTodo(w, r, ps)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
