package main

import "densho/api"

func main() {
	api := api.NewApi(":3000")

	api.Start()
}
