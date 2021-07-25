package database

import (
	"database/sql"
	"time"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:redmond@tcp(127.0.0.1:3306)/inventorydb")
	// if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}