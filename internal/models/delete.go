package models

import (
	"fmt"
	"net/http"

	"github.com/ihor-matiushenko/todo-list/configs"
	"github.com/julienschmidt/httprouter"
)

// DeleteTodo ...
func DeleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := configs.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
