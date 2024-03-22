package db

import (
	"database/sql"
)

type PostgresDBData struct {
	Username      string
	Password      string
	Database_ip   string
	Database_name string
}

type Data struct {
	Id          int    `field:"id"`
	Title       string `field:"title"`
	Description string `field:"description"`
}

// DBClient defines the interface for database operations
type DBClient interface {
	Select() ([]Data, error)
	Insert(title, description string) (int, error)
	GetDataByID(id int) (*Data, error)
	DeleteByID(id int) (bool, error)
}

// SQLDBClient represents a concrete implementation of the DBClient interface
type SQLDBClient struct {
	Conn *sql.DB
}
