package main

import (
	"github.com/JvstRyan/go-event-rest-api/db"
	"github.com/JvstRyan/go-event-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	godotenv.Load()
	server.Run(":8080")
}
