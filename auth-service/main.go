package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=postgres user=postgres password=postgres dbname=auth_db sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer db.Close()
	fmt.Println("Connected to PostgreSQL in auth-service")
}
