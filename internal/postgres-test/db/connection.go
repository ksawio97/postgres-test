package db

import (
	"database/sql"
	"fmt"
)

func ConnectToDB(data PostgresDBData) (*sql.DB, error) {
	// Connect to database
	conn, err := sql.Open("postgres", data.connectionStringBuilder())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewSQLDBClient creates a new SQLDBClient instance
func NewSQLDBClient(conn *sql.DB) *SQLDBClient {
	return &SQLDBClient{Conn: conn}
}

func (postgresDBData PostgresDBData) connectionStringBuilder() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		postgresDBData.Username,
		postgresDBData.Password,
		postgresDBData.Database_ip,
		postgresDBData.Database_name)
}
