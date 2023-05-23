package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var MysqlDb *sql.DB = MysqlDbInstance()

func MysqlDbInstance() *sql.DB {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlUrl := os.Getenv("MYSQL_URL")

	db, err := sql.Open("mysql", mysqlUrl)

	if err != nil {
		panic(err)
	}

	return db
}
