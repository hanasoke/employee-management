package main

import (
	"employee-management/database"
	"employee-management/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.InitDB()

	// Initialize router
	employeeHandler := handlers.NewEmployeeHandler()

	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Routes
	api := router.Group("/api/v1")
	{
		employees := api.Group("/employees")
		{
			employees.GET("", employeeHandler.GetEmployee)
			employees.GET("/stats", employeeHandler.GetEmployeeStats)

		}
	}
}
