package api

import (
	"Hrms/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	route.Get("/", controllers.Index)

	employee := a.Group("/api/v1/employee")
	employee.Get("/", controllers.GetEmployees)
	employee.Get("/:id", controllers.GetEmployee)
	employee.Post("/", controllers.CreateEmployee)
	employee.Patch("/:id", controllers.UpdateEmployee)
	employee.Delete("/:id", controllers.DeleteEmployee)

}
