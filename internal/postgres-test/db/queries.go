package db

import (
	"database/sql"
	"reflect"
)

func Select(conn *sql.DB) ([]Data, error) {
	// Select data from Data table
	rows, err := conn.Query(`SELECT id, title, description FROM public."Data";`)
	if err != nil {
		return nil, err
	}
	elements := []Data{}

	for rows.Next() {
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
		// Load query values into parameters
		rows.Scan(cols...)

		elements = append(elements, element)
	}
	return elements, nil
}
