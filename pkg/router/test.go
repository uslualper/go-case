package router

import (
	"github.com/gofiber/fiber/v2"

	handlersV1 "yuka-case/pkg/handlers/api/v1"
)

type Test struct{}

func (t *Test) SetupRoutes(r fiber.Router) {
	v1 := r.Group("/v1")

	test := v1.Group("/test")

	handler := handlersV1.Test{}

	test.Post("/", handler.Test)
}
