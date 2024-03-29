package main

import (
	"log"
	"postgres-test/test/api"
	"postgres-test/test/internal/postgres-test/db"
	"postgres-test/test/internal/postgres-test/flags"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	postgresDBData := flags.ReadPostgresDBDataFlags()

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
