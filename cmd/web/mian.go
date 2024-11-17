package main

import (
	"bitcoin-rate/db"
	"bitcoin-rate/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server := gin.Default()

	db.InitPgRepository()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
