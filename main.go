package main

import (
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Printf("--------------- %s Started------------", os.Getenv("APP_NAME"))
	app := gin.Default()

	app.Static("/images/", "./images")

	app.GET("/", func(c *gin.Context) {
		log.Printf("------------- %s function called ---------------", os.Getenv("APP_NAME"))
		c.File(path.Join(".", "index.html"))
	})

	app.Run(":3000")
	log.Printf("Server started of %v", 3000)
}
