package routes

import (
	"github.com/gofiber/fiber/v2"
)

// NotFoundRoute function voor alle 404 errors
func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			// Stuur een 404 status en JSON response met de reden
			// @TODO: Vervang de JSON response met een mooie 404-pagina
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "Sorry, deze endpoint kon niet gevonden worden.",
			})
		},
	)
}
