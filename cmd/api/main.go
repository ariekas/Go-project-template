package main

import (
	"fmt"
	"template-golang/internal/config"
	"template-golang/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := config.Database()
	if err != nil {
		fmt.Printf("Failed to connect database, Error: %s", err)
		return
	}
	defer db.Close()

	app := routes.Routes(db)

	if err := app.Listen(":8080"); err != nil {
		fmt.Printf("Failed to start server, Error: %s", err)
	}
}
