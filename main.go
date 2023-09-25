package main

import (
	"log"
	"os"

	"github.com/RahulMj21/affine-clone/config"
	"github.com/RahulMj21/affine-clone/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	config.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	engine := handlebars.New("./views", ".hbs")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(logger.New())

	app.Static("/", "./static")

	routes.PageRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
