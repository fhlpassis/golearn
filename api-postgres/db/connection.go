package db

import (
	"database/sql"
	"fmt"

	"github.com/fhlpassis/golang-postgres/configs"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
	db, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
