package main

import (
	"github.com/Omar-Walid/shape-calc/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/shapes/calculate", handlers.CalculateShape)
	}

	router.Run(":8080")
}
