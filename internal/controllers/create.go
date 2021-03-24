package controllers

import (
	"net/http"

	// ...
	_ "github.com/lib/pq"

	"github.com/ihor-matiushenko/todo-list/internal/models"
	"github.com/julienschmidt/httprouter"
)

// CreateTodo ...
func CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	models.CreateTodo(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
