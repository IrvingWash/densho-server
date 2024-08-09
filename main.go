package main

import (
	"densho/api"
	"densho/db"
)

func main() {
	db := db.NewDb("./testdb")

	api := api.NewApi(":3000", &db)

	api.Start()

	db.Close()
}
