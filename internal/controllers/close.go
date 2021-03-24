package controllers

import (
	"net/http"

	// ...
	_ "github.com/lib/pq"

	"github.com/ihor-matiushenko/todo-list/internal/models"
	"github.com/julienschmidt/httprouter"
)

// CloseTodo ...
func CloseTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	models.CloseTodo(w, r, ps)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
