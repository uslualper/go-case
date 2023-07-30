package router

import (
	"github.com/gofiber/fiber/v2"

	_ "yuka-case/docs"
)

type Router interface {
	SetupRoutes(r fiber.Router)
}

func SetupRoutes(a *fiber.App) {

	api := a.Group("/api")

	routes := []Router{
		&System{},
		&Test{},
		&Swagger{},
		&Location{},
		&Routing{},
	}

	for _, route := range routes {
		route.SetupRoutes(api)
	}

}
