package main

import (
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/gin-gonic/gin"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/router"


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
	gin := gin.Default()

	// Set up routes
	router.Setup(env, timeout, db, gin)

	// Start the server
	gin.Run(env.LocalServerPort)
}
