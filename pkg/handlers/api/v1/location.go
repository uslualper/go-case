package handlers

import (
	"github.com/gofiber/fiber/v2"

	payloadSchema "yuka-case/pkg/schema/payload/v1"
	"yuka-case/pkg/services"
	"yuka-case/pkg/utils/http"
	"yuka-case/pkg/utils/validate"
)

type Location struct{}

// Location
// @Summary Get Locations
// @Description Get locations
// @Tags Location
// @Produce json
// @Param page query int false "page" default(1)
// @Router /v1/location [get]
func (l *Location) Get(c *fiber.Ctx) error {
	paginate := http.InitPaginate(c)

	locations, err := (&services.Location{}).GetLocations(paginate.Page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(locations)
}

// Location
// @Summary Get Location
// @Description Get Location
// @Tags Location
// @Produce json
// @Param id path int true "Location ID"
// @Router /v1/location/{id} [get]
func (l *Location) Find(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	location, err := (&services.Location{}).GetLocationById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(location)
}

// Location
// @Summary Crate Location
// @Description Create Location
// @Tags Location
// @Accept json
// @Produce json
// @Param payload body payloadSchema.Location true "Location Payload"
// @Router /v1/location [post]
func (l *Location) Create(c *fiber.Ctx) error {
	payload := new(payloadSchema.Location)
	c.BodyParser(payload)
	if err := validate.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	if err := (&services.Location{}).CreateLocation(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		return c.SendStatus(fiber.StatusOK)
	}
}

// Location
// @Summary Update Location
// @Description Update Location
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Param payload body payloadSchema.Location true "Location Payload"
// @Router /v1/location/{id} [put]
func (l *Location) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	payload := new(payloadSchema.Location)
	c.BodyParser(payload)
	if err := validate.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	if err := (&services.Location{}).UpdateLocation(id, payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		return c.SendStatus(fiber.StatusOK)
	}
}

// Location
// @Summary Delete Location
// @Description Delete Location
// @Tags Location
// @Produce json
// @Param id path int true "Location ID"
// @Router /v1/location/{id} [delete]
func (l *Location) Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := (&services.Location{}).DeleteLocation(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		return c.SendStatus(fiber.StatusOK)
	}
}
