package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetBD() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "synjcruz"
	dbPass := "Misuperllave2020#"
	dbName := "employees"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
