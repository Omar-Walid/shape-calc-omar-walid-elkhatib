package handlers

import (
	"math"
	"net/http"

	"github.com/Omar-Walid/shape-calc/models"
	"github.com/Omar-Walid/shape-calc/services"
	"github.com/gin-gonic/gin"
)

func CalculateShape(c *gin.Context) {
	var input models.ShapeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	shape, err := services.CreateShape(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.ShapeResponse{
		Shape:     input.Type,
		Area:      round(shape.Area(), 2),
		Perimeter: round(shape.Perimeter(), 2),
	}

	c.JSON(http.StatusOK, response)
}

func round(val float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(val*pow) / pow
}
