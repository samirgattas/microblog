package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(userDB string, passwordDB string, addressDB string, portDB int64) (*sql.DB, error) {
    nameDB := "microblog"
	sql, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userDB, passwordDB, addressDB, portDB, nameDB))
	if err != nil {
		return nil, err
	}
	return sql, nil
}
