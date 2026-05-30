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
	FullName   string      `db:"full_name" json:"full_name"`
	LeaveType  LeaveType   `db:"leave_type" json:"leave_type"`
	StartDate  time.Time   `db:"start_date" json:"start_date"`
	EndDate    time.Time   `db:"end_date" json:"end_date"`
	Reason     string      `db:"reason" json:"reason"`
	Status     LeaveStatus `db:"status" json:"status"`
	ReviewerID *uuid.UUID  `db:"reviewer_id" json:"reviewer_id,omitempty"`
	CreatedAt  time.Time   `db:"created_at" json:"created_at"`
}
