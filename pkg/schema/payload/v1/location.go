package v1

type Location struct {
	Name        string  `json:"name" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
	MarkerColor string  `json:"marker"`
}
