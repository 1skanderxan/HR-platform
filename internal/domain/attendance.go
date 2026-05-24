package domain

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
)

type Attendance struct {
	ID         uuid.UUID        `db:"id" json:"id"`
	EmployeeID uuid.UUID        `db:"employee_id" json:"employee_id"`
	FullName   string           `db:"full_name" json:"full_name"`
	CheckIn    *time.Time       `db:"check_in" json:"check_in"`
	CheckOut   *time.Time       `db:"check_out" json:"check_out"`
	Status     AttendanceStatus `db:"status" json:"status"`
	Latitude   float64          `db:"latitude" json:"latitude"`
	Longitude  float64          `db:"longitude" json:"longitude"`
	Date       time.Time        `db:"date" json:"date"`
}
