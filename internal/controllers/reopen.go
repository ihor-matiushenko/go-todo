package controllers

import (
	"net/http"

	// ...
	_ "github.com/lib/pq"

	"github.com/ihor-matiushenko/todo-list/internal/models"
	"github.com/julienschmidt/httprouter"
)

// ReopenTodo ...
func ReopenTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	models.ReopenTodo(w, r, ps)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
