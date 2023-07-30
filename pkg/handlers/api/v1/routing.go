package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"yuka-case/pkg/services"
)

type Routing struct{}

// Routing
// @Summary Get Routing
// @Description Get Routing
// @Tags Routing
// @Produce json
// @Param id path int true "Location ID"
// @Router /v1/routing/{id} [get]
func (rg *Routing) Fetch(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	locationService := services.Location{}
	locations, _ := locationService.GetAllLocations()
	targetLocation, _ := locationService.GetLocationById(id)

	routingService := services.Routing{}

	type LocationDistance struct {
		Name             string  `json:"name"`
		Latitude         float64 `json:"latitude"`
		Longitude        float64 `json:"longitude"`
		DistanceToTarget float64 `json:"distance_to_target"`
	}

	var locationsDistance []LocationDistance
	for _, location := range locations {
		locationsDistance = append(locationsDistance, LocationDistance{
			Name:             location.Name,
			Latitude:         location.Latitude,
			Longitude:        location.Longitude,
			DistanceToTarget: routingService.Distance(location.Latitude, location.Longitude, targetLocation.Latitude, targetLocation.Longitude),
		})
	}

	return c.JSON(locationsDistance)
}
