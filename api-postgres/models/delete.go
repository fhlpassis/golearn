package models

import (
	"github.com/fhlpassis/golang-postgres/db"
)

func Delete(id int64) (int64, error) {
	db, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sql := `DELETE FROM todos WHERE id = $1`

	res, err := db.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
