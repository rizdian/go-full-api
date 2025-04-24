package main

import (
	"github.com/rizdian/go-full-api/config"
	"github.com/rizdian/go-full-api/internal/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := config.ConnectDatabase()
	r := gin.Default()

	api := r.Group("/api")
	{
		handler.RegisterAuthRoutes(api, db)
		handler.RegisterUserRoutes(api, db)
		handler.RegisterProductRoutes(api, db)
		handler.RegisterOrderRoutes(api, db)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
