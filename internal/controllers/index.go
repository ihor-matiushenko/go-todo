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
