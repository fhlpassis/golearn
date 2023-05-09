package models

import (
	"github.com/fhlpassis/golang-postgres/db"
)

func Insert(todo Todo) (id int, err error) {
	db, err := db.Connect()
	if err != nil {
		return
	}
	defer db.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
