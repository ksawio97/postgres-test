package main

import (
	"log"
	"postgres-test/test/api"
	"postgres-test/test/internal/postgres-test/db"

	"github.com/gin-gonic/gin"
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

	client := db.NewSQLDBClient(conn)
	handler := api.NewHandler(client)

	r := gin.Default()
	handler.SetupRoutes(r)
}
