package flags

import (
	"flag"
	"postgres-test/test/internal/postgres-test/db"
)

func ReadPostgresDBDataFlags() db.PostgresDBData {
	username := flag.String("u", "postgres", "Username for postgres connection")
	password := flag.String("p", "postgres", "Password for postgres connection")
	db_ip := flag.String("h", "127.0.0.1", "Host (ip) for postgres connection")
	db_name := flag.String("db", "test", "Database name for postgres connection")

	flag.Parse()

	return db.PostgresDBData{
		Username:      *username,
		Password:      *password,
		Database_ip:   *db_ip,
		Database_name: *db_name,
	}
}
