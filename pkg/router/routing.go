package router

import (
	"github.com/gofiber/fiber/v2"

	handlersV1 "yuka-case/pkg/handlers/api/v1"
)

type Routing struct{}

func (rg *Routing) SetupRoutes(r fiber.Router) {
	v1 := r.Group("/v1")

	location := v1.Group("/routing")

	handler := handlersV1.Routing{}

	location.Get("/:id", handler.Fetch)
}
