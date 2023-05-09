package models

import (
	"github.com/fhlpassis/golang-postgres/db"
)

func Get(id int64) (todo Todo, err error) {
	db, err := db.Connect()
	if err != nil {
		return
	}
	defer db.Close()

	sql := `SELECT id, title, description, done FROM todos WHERE id = $1`

	err = db.QueryRow(sql, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
