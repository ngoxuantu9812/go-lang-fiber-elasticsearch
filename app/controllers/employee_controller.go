package controllers

import (
	"Hrms/app/models"
	"Hrms/pkg/logging"
	"Hrms/pkg/utils"
	"Hrms/platform/database"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func GetEmployees(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all employees.
	employees, err := db.GetEmployees()

	if err != nil {
		// Return, if employees not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":     true,
			"msg":       "employees were not found",
			"count":     0,
			"employees": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"count":     len(employees),
		"employees": employees,
	})
}

func GetEmployee(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	db, err := database.OpenDBConnection()
	employee, err := db.GetEmployee(id)
	if err != nil {
		// Return, if employees not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":     true,
			"msg":       "employees were not found",
			"count":     0,
			"employees": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"employee": employee,
	})
}

func CreateEmployee(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	employee := &models.Employee{}
	employee.ID = uuid.New()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	if err := c.BodyParser(employee); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fmt.Println(employee)
	validate := utils.NewValidator()

	if err := validate.Struct(employee); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create book by given model.
	if err := db.CreateEmployee(employee); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	employeeJson, err := json.Marshal(employee)
	if err != nil {
		logging.CreateLog("Create Employee Success", logging.LevelInfo, employeeJson)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	logging.CreateLog("Create Employee Success", logging.LevelInfo, employeeJson)

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"employee": employee,
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	id, err := uuid.Parse(c.Params("id"))
	employee := &models.Employee{}
	employee.ID = id
	employee.UpdatedAt = time.Now()
	if err := c.BodyParser(employee); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedEmployee, err := db.GetEmployee(id)
	if err != nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with this ID not found",
		})
	}
	validate := utils.NewValidator()

	if err := validate.Struct(employee); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update employee by given model.
	if err := db.UpdateEmployee(foundedEmployee.ID, employee); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"employee": employee,
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	id, err := uuid.Parse(c.Params("id"))

	if err := db.DeleteEmployee(id); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Delete employee success",
	})
}
