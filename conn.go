package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/dvdrental", os.Getenv("PG_USER"), os.Getenv("PG_PASS")))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	rows, err := db.Query("SELECT id, name FROM public.customer_list WHERE id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
