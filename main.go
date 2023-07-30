package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"

	"yuka-case/pkg/config"
	"yuka-case/pkg/db"
	"yuka-case/pkg/i18n"
	"yuka-case/pkg/router"
)

// @title API
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @contact.name   Alper Uslu
// @contact.email  uslualper@outlook.com

// @BasePath /api
func main() {

	port := config.Config("APP_PORT")

	app := fiber.New(fiber.Config{
		Immutable: true,
		Views:     html.New("./views", ".html"),
	})

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	allowOrigins := []string{}

	if config.Config("APP_DEBUG") == "true" {
		allowOrigins = append(allowOrigins, "*")
	} else {
		allowOrigins = append(allowOrigins, config.Config("APP_URL"))
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(allowOrigins, ","),
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	db.SetupDB()

	router.SetupRoutes(app)

	i18n.SetupI18n()

	app.Listen(":" + port)
}
