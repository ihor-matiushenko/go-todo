package models

// Todo ...
type Todo struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Done  bool   `json:"done" db:"done"`
}
