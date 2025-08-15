package services

import (
	"errors"
	"github.com/Omar-Walid/shape-calc/models"
	"strings"
)

func CreateShape(input models.ShapeRequest) (models.Shape, error) {
	shapeType := strings.ToLower(input.Type)

	switch shapeType {
	case "circle":
		if input.Radius <= 0 {
			return nil, errors.New("radius must be positive")
		}
		return models.Circle{Radius: input.Radius}, nil

	case "rectangle":
		if input.Length <= 0 || input.Width <= 0 {
			return nil, errors.New("length and width must be positive")
		}
		return models.Rectangle{Length: input.Length, Width: input.Width}, nil

	case "triangle":
		if input.Base <= 0 || input.Height <= 0 || input.Side1 <= 0 || input.Side2 <= 0 || input.Side3 <= 0 {
			return nil, errors.New("all triangle sides and height must be positive")
		}
		if !isValidTriangle(input.Side1, input.Side2, input.Side3) {
			return nil, errors.New("triangle inequality violated")
		}
		return models.Triangle{
			Base:   input.Base,
			Height: input.Height,
			Side1:  input.Side1,
			Side2:  input.Side2,
			Side3:  input.Side3,
		}, nil

	default:
		return nil, errors.New("invalid shape type")
	}
}

func isValidTriangle(a, b, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}
