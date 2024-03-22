package db

import (
	"database/sql"
	"fmt"
	"reflect"
)

func Select(conn *sql.DB) ([]Data, error) {
	// Select data from Data table
	rows, err := conn.Query(`SELECT id, title, description FROM public."Data";`)
	if err != nil {
		return nil, err
	}
	elements := []Data{}

	var element Data

	// Get pointers to every value in Data type
	s := reflect.ValueOf(&element).Elem()

	// Get values count
	numCols := s.NumField()

	// Gather pointers to this fields
	cols := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		cols[i] = field.Addr().Interface()
	}

	for rows.Next() {
		// Load query values into parameters
		rows.Scan(cols...)

		elements = append(elements, element.Copy())
	}
	return elements, nil
}

func Insert(conn *sql.DB, title, description string) (int, error) {
	var id int
	result := conn.QueryRow(fmt.Sprintf(`INSERT INTO public."Data"(title, description) VALUES ('%s', '%s') RETURNING id;`,
		title,
		description))
	if result.Err() != nil {
		return 0, result.Err()
	}
	result.Scan(&id)
	return id, nil
}
