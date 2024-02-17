package web

import (
	"Hrms/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("")
	route.Get("/", controllers.Index)
}
