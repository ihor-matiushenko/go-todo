package configs

import (
	"database/sql"
	"fmt"

	// ...
	_ "github.com/lib/pq"
)

// DB ...
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "dbname=todo_list sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database")
}
