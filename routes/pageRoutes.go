package routes

import "github.com/gofiber/fiber/v2"

func PageRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Affine",
		})
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("dashboard", fiber.Map{
			"Title": "Affine - Dashboard",
		})
	})

	app.Get("/pricing", func(c *fiber.Ctx) error {
		return c.Render("pricing", fiber.Map{
			"Title": "Affine - Pricing",
		})
	})

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("docs", fiber.Map{
			"Title": "Affine - Docs",
		})
	})

	app.Get("/support", func(c *fiber.Ctx) error {
		return c.Render("support", fiber.Map{
			"Title": "Affine - Support",
		})
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{
			"Title": "Affine - Contact",
		})
	})

}
