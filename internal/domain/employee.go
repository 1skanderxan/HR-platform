package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleEmployee Role = "employee"
	RoleHR       Role = "hr"
	RoleAdmin    Role = "admin"
)

type Employee struct {
	ID           uuid.UUID `db:"id" json:"id"`
	FullName     string    `db:"full_name" json:"full_name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Role         Role      `db:"role" json:"role"`
	DepartmentID uuid.UUID `db:"department_id" json:"department_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
