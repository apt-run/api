package main

import (
	"fmt"
	"main/database"
	"main/routes"

	"github.com/sudo-adduser-jordan/gcolor"
)

func main() {

	fmt.Println()
	fmt.Print(gcolor.BlueLabel(" Go 1.22 "))
	fmt.Print(gcolor.PurpleLabel(" Postgres 15 "))
	fmt.Println(gcolor.BlueLabel(" pgx v5 "))
	fmt.Println()

	database.ConnectToDatabase()
	database.MigrateDatabase()
	routes.SetupRoutes()
}
