package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Listed GO Todo",
	})
}