package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pyaqubzade/knative/client"
	"github.com/pyaqubzade/knative/config"
	"github.com/pyaqubzade/knative/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(recover.New())

	api := app.Group("/api")
	c := client.NewClient()
	handler.NewHealthHandler(app)
	handler.NewHandler(api, c)

	port := config.Props.Port
	log.Info("Starting server at port: ", port)
	log.Fatal(app.Listen(":" + port))
}
