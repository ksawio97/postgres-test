package db

import "fmt"

func (data Data) String() string {
	return fmt.Sprintf("Id: %s Title: %s Description: %s", fmt.Sprint(data.Id), data.Title, data.Description)
}

func (data Data) Copy() Data {
	return Data{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
	}
}
