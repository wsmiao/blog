package service

import (
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

// ConnectDB connects to the database.
func ConnectDB() {
	var err error
	db, err = sqlx.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatalln(err)
	}
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Fatalln("Disconnect from database failed: " + err.Error())
	}
}
