package models

import (
	"fmt"
	"net/http"

	"github.com/ihor-matiushenko/todo-list/configs"
	"github.com/julienschmidt/httprouter"
)

// CloseTodo ...
func CloseTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := configs.DB.Exec("UPDATE todos SET done = $1 WHERE id = $2", true, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}
