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

	id, err := db.Insert(conn, "insert", "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted row with auto-assigned id: " + fmt.Sprint(id))

	// Read all rows in database
	elements, err := db.Select(conn)
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range elements {
		fmt.Println(element.String())
	}
}
