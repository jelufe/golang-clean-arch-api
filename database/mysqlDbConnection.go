package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB = MysqlDbInstance()

func MysqlDbInstance() *sql.DB {
	loadEnv()

	mysqlUrl := os.Getenv("MYSQL_URL")

	db, err := sql.Open("mysql", mysqlUrl)

	if err != nil {
		panic(err)
	}

	return db
}
