package entity

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MarkerColor string  `json:"marker"`
}
