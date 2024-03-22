package main

import (
	"fmt"
	"log"
	"postgres-test/test/internal/postgres-test/db"

	_ "github.com/lib/pq"
)

func main() {
	postgresDBData := db.PostgresDBData{
		Username:      "postgres",
		Password:      "postgres",
		Database_ip:   "127.0.0.1",
		Database_name: "test",
	}

	// Connect to database
	conn, err := db.ConnectToDB(postgresDBData)
	if err != nil {
		log.Fatal(err)
	}

	// Read all rows in database
	elements, err := db.Select(conn)
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range elements {
		fmt.Println(element.String())
	}
}
