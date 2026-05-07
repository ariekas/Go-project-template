package main

import (
	"fmt"
	"template-golang/internal/config"
	"template-golang/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()

	godotenv.Load()

	db, err := config.Database()
	if err != nil {
		fmt.Printf("Failed to connect database, Error: %s", err)
		return
	}

	routes.Routes(db)

	app.Run(":8080")
}
