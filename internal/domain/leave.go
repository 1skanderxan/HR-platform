package domain

import (
	"time"

	"github.com/google/uuid"
)

type LeaveType string
type LeaveStatus string

const (
	LeaveAnnual LeaveType = "annual"
	LeaveSick   LeaveType = "sick"
	LeaveUnpaid LeaveType = "unpaid"

	LeaveStatusPending  LeaveStatus = "pending"
	LeaveStatusApproved LeaveStatus = "approved"
	LeaveStatusRejected LeaveStatus = "rejected"
)

type LeaveRequest struct {
	ID         uuid.UUID   `db:"id" json:"id"`
	EmployeeID uuid.UUID   `db:"employee_id" json:"employee_id"`
	LeaveType  LeaveType   `db:"leave_type" json:"leave_type"`
	StartDate  time.Time   `db:"start_date" json:"start_date"`
	EndDate    time.Time   `db:"end_date" json:"end_date"`
	Reason     string      `db:"reason" json:"reason"`
	Status     LeaveStatus `db:"status" json:"status"`
	ReviewedBy *uuid.UUID  `db:"reviewed_by" json:"reviewed_by,omitempty"`
	CreatedAt  time.Time   `db:"created_at" json:"created_at"`
}
