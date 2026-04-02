package main

import (
	"fmt"
	"log"
	"template-golang/internal/config"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	/*
	routes "Golang/internal/routes" --> import modul routes app
	*/
)

func main() {
	app := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment")
	}

	_, err := config.ConnectionDB()
	// db, err := config.ConnectionDB() --> Deklarasi database connection


	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}

	app.GET("/", func(c *gin.Context) {
		c.String(200, "Back IS Running")
	})

	/*
	app := routes.New(db) --> Menjalankan database connection
	*/

	app.Run(":8080")
}
