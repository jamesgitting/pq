package main

import (
	"database/sql"
	"github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "user=james dbname=tiro sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

}
