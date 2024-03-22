package db

import (
	"fmt"
	"reflect"
)

func (c *SQLDBClient) Select() ([]Data, error) {
	// Select data from Data table
	rows, err := c.Conn.Query(`SELECT id, title, description FROM public."Data";`)
	if err != nil {
		return nil, err
	}
	elements := []Data{}

	var element Data
	cols := getPointersToCols(element)

	for rows.Next() {
		// Load query values into parameters
		rows.Scan(cols...)

		elements = append(elements, element.Copy())
	}
	return elements, nil
}

func (c *SQLDBClient) Insert(title, description string) (int, error) {
	var id int
	result := c.Conn.QueryRow(fmt.Sprintf(`INSERT INTO public."Data"(title, description) VALUES ('%s', '%s') RETURNING id;`,
		title,
		description))
	if result.Err() != nil {
		return 0, result.Err()
	}

	result.Scan(&id)
	return id, nil
}

func (c *SQLDBClient) GetDataById(id int) (*Data, error) {
	row := c.Conn.QueryRow(fmt.Sprintf(`SELECT id, title, description FROM public."Data" WHERE id = %s;`, fmt.Sprint(id)))

	if row.Err() != nil {
		return nil, row.Err()
	}

	var element Data
	cols := getPointersToCols(element)

	row.Scan(cols...)
	return &element, nil
}

// returns true if row has been deleted
func (c *SQLDBClient) DeleteById(id int) (bool, error) {
	result, err := c.Conn.Exec(fmt.Sprintf(`DELETE FROM public."Data" WHERE id = %s;`, fmt.Sprint(id)))
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected == 1, nil
}

func getPointersToCols(element Data) []interface{} {
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

	return cols
}
