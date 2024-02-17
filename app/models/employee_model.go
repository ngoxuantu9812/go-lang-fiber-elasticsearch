package models

import (
	"time"

	"github.com/google/uuid"
)

// Employee struct to describe employee object.
type Employee struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,lte=255"`
	FirstName    string    `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName     string    `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Position     string    `db:"position" json:"position" validate:"required,lte=255"`
	Salary       int32     `db:"salary" json:"salary" validate:"required"`
	DepartmentId int32     `db:"department_id" json:"department_id"`
	Password     string    `db:"password" json:"password" validate:"required"`
}
