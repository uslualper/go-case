package router

import (
	"github.com/gofiber/fiber/v2"

	handlersV1 "yuka-case/pkg/handlers/api/v1"
)

type Location struct{}

func (l *Location) SetupRoutes(r fiber.Router) {
	v1 := r.Group("/v1")

	location := v1.Group("/location")

	handler := handlersV1.Location{}

	location.Get("/", handler.Get)
	location.Get("/:id", handler.Find)
	location.Post("/", handler.Create)
	location.Put("/:id", handler.Update)
	location.Delete("/:id", handler.Delete)
}
