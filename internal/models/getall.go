package models

import (
	"fmt"

	"github.com/ihor-matiushenko/todo-list/configs"
)

// GetTodos ...
func GetTodos() ([]Todo, error) {
	rows, err := configs.DB.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			fmt.Println(err)
			continue
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
