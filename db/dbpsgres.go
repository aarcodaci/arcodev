package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "pgdb-arcodev-19569.nodechef.com"
	port     = 2427
	user     = "ncuser_15634"
	password = "LLsJrNSYCgFjTDT9exGBXYKm1uhOnB"
	dbname   = "arcodev"
)

func DbConnect() *sql.DB {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	return db
}

func DoPostgress() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
