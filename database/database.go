package database

import (
	"database/sql"

	"fmt"
	"time"

	//mysql-driver imported for effect
	_ "github.com/go-sql-driver/mysql"
)

//Init sets up the database
func Init() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tonaira")
	if err != nil {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		fmt.Println("Ping Failed")
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(5)
	dbConn.SetConnMaxLifetime(time.Second * 10)

	return dbConn
}
