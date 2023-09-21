package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "libary"
	password = "secret"
	dbname   = "libary"
)

var db *sql.DB

func init() {
	// Connect to the database during package initialization
	log.Printf(`Started: creating database connection for host=%s port=%d user=%s password=XXXXXX dbname=%s`,
		host, port, user, dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf(`End: created database connection for host=%s port=%d user=%s password=XXXXXX dbname=%s`,
		host, port, user, dbname)
}
