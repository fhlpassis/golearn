package models

import (
	"github.com/fhlpassis/golang-postgres/db"
)

func Update(id int64, todo Todo) (int64, error) {
	db, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sql := `UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4`

	res, err := db.Exec(sql, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
