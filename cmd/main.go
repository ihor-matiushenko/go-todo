package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ihor-matiushenko/todo-list/internal/router"
)

func main() {
	fmt.Println("Serving...")

	log.Fatal(http.ListenAndServe(":8080", router.New()))
}
