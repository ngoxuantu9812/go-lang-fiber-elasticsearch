package controllers

import (
	"Hrms/pkg/logging"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	logging.CreateLog("Helloo, I'm here hehehe3333", logging.LevelInfo)
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  "Hello World",
	})
}
