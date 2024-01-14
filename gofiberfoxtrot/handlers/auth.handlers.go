package handlers

import (
	"github.com/gofiber/fiber/v2"
)


func HandleViewHome(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
