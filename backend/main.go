package main

import (
	"ucc-arq-soft-I/app"
	"ucc-arq-soft-I/database"
)

func main() {
	database.Connect()
	database.Migrate()

	app.StartRoute()
}
