package queries

import (
	"Hrms/app/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// EmployeeQueries struct for queries from Employee model.
type EmployeeQueries struct {
	*sqlx.DB
}

// GetEmployees method for getting all employees.
func (q *EmployeeQueries) GetEmployees() ([]models.Employee, error) {
	// Define employees variable.
	employees := []models.Employee{}

	// Define query string.
	query := `SELECT * FROM employees`

	// Send query to database.
	err := q.Select(&employees, query)
	fmt.Println(err)
	if err != nil {
		// Return empty object and error.
		return employees, err
	}

	// Return query result.
	return employees, nil
}

// GetEmployee method for getting one employee by given ID.
func (q *EmployeeQueries) GetEmployee(id uuid.UUID) (models.Employee, error) {
	// Define employee variable.
	employee := models.Employee{}

	// Define query string.
	query := `SELECT * FROM employees WHERE id = $1`

	// Send query to database.
	err := q.Get(&employee, query, id)
	if err != nil {
		// Return empty object and error.
		return employee, err
	}

	// Return query result.
	return employee, nil
}

// CreateEmployee method for creating employee by given Employee object.
func (q *EmployeeQueries) CreateEmployee(e *models.Employee) error {
	// Define query string.
	query := `INSERT INTO employees VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// Send query to database.
	_, err := q.Exec(query, e.ID, e.CreatedAt, e.UpdatedAt, e.Email, e.FirstName, e.LastName, e.Position, e.Salary, e.DepartmentId, e.Password)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateEmployee method for updating employee by given Employee object.
func (q *EmployeeQueries) UpdateEmployee(id uuid.UUID, b *models.Employee) error {
	// Define query string.
	query := `UPDATE employees SET updated_at = $2, email = $3, first_name = $4, last_name = $5, position = $6, salary = $7, department_id = $8, password = $9 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.UpdatedAt, b.Email, b.FirstName, b.LastName, b.Position, b.Salary, b.DepartmentId, b.Password)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteEmployee method for delete employee by given ID.
func (q *EmployeeQueries) DeleteEmployee(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM employees WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
