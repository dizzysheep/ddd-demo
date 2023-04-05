package main

import (
	"ddd-demo/app/interfaces/router"
	"ddd-demo/infrastructure/database"
)

func main() {
	db := database.NewDB()

	r := router.SetupRouter(db)
	r.Run(":8080")
}
