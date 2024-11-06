package main

import (
	"time"
	"log"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	env := config.NewEnv()

	// Connect to MongoDB
	dbClient := config.ConnectMongoDB(env)
	db := dbClient.Database(env.DBName)

	// Set up server timeout
	timeout := time.Duration(env.ContextTimeout) * time.Second

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,  // Set to true if your frontend sends credentials (e.g., cookies, Authorization headers)
		MaxAge:           12 * time.Hour,
	}))

	// Set up routes
	router.Setup(env, timeout, db, r)

	// Start the server
	if err := r.Run(env.LocalServerPort); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
