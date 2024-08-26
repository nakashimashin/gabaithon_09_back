package main

import (
	"gabaithon-09-back/database"
	"gabaithon-09-back/routes"
)

func main() {
	database.Migrate()
	router := routes.GetApiRouter(database.DB)
	router.Run(":8082")
}
