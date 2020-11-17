package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectSqlxDatabase() (dbx *sqlx.DB) {

	dbx, err := sqlx.Open("mysql", "root:admin@123@tcp(localhost:3306)/activity_tracker_db")
	err = dbx.Ping()

	if err != nil {
		log.Fatalln(err)

	} else {
		log.Println("Database Sqlx Connected")

	}
	//defer dbx.Close()
	return dbx
}
