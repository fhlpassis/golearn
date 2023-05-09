package models

import (
	"github.com/fhlpassis/golang-postgres/db"
)

func GetAll() (todos []Todo, err error) {
	db, err := db.Connect()
	if err != nil {
		return
	}
	defer db.Close()

	sql := `SELECT id, title, description, done FROM todos`

	rows, err := db.Query(sql)

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return
}
