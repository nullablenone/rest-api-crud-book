package main

import (
	"rest-api-crud-books/config"
	"rest-api-crud-books/models"
)

func main(){

	config.ConnectDB()
	models.Migrate(config.DB)
}