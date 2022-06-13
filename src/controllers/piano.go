package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Piano(c *fiber.Ctx) error {
	return c.Send([]byte("test"))
}
