package models

import (
	"fmt"
	"net/http"

	"github.com/ihor-matiushenko/todo-list/configs"
	"github.com/julienschmidt/httprouter"
)

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
