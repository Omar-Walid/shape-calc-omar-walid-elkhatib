package models

type ShapeRequest struct {
	Type   string  `json:"type"`
	Radius float64 `json:"radius,omitempty"`
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Base   float64 `json:"base,omitempty"`
	Height float64 `json:"height,omitempty"`
	Side1  float64 `json:"side1,omitempty"`
	Side2  float64 `json:"side2,omitempty"`
	Side3  float64 `json:"side3,omitempty"`
}

type ShapeResponse struct {
	Shape     string  `json:"shape"`
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`
}
