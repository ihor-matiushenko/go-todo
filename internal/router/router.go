package router

import (
	"net/http"

	"github.com/ihor-matiushenko/todo-list/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

// New ...
func New() *httprouter.Router {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("../web/static"))

	router.GET("/", controllers.Index)
	router.POST("/create/", controllers.CreateTodo)
	router.POST("/todos/:id/update/", controllers.UpdateTodo)
	router.GET("/todos/:id/close/", controllers.CloseTodo)
	router.GET("/todos/:id/reopen/", controllers.ReopenTodo)
	router.GET("/todos/:id/delete/", controllers.DeleteTodo)

	return router
}
