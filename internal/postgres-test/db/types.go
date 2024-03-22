package db

import "fmt"

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
