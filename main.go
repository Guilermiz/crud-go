package main

import (
	"github.com/Guilermiz/crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	group := router.Group("/user")
	routes.InitUserRoutes(group)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
