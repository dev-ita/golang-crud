package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, pass, db)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
