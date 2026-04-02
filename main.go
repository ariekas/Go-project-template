package main

import "github.com/gin-gonic/gin"

func main(){
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.String(200, "Back IS Running")
	})

	app.Run(":8080")
}