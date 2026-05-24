package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type EmployeeRepository interface {
	Create(ctx context.Context, emp *Employee) error
	GetByID(ctx context.Context, id uuid.UUID) (*Employee, error)
	GetByEmail(ctx context.Context, email string) (*Employee, error)
	Update(ctx context.Context, emp *Employee) error
	List(ctx context.Context) ([]*Employee, error)
}

type AttendanceRepository interface {
	CheckIn(ctx context.Context, att *Attendance) error
	CheckOut(ctx context.Context, employeeID uuid.UUID) error
	GetTodayByEmployee(ctx context.Context, employeeID uuid.UUID) (*Attendance, error)
	ListByEmployee(ctx context.Context, employeeID uuid.UUID) ([]*Attendance, error)
	ListAll(ctx context.Context) ([]*Attendance, error)
	ListByEmployeePeriod(ctx context.Context, employeeID uuid.UUID, from, to time.Time) ([]*Attendance, error)
}

type LeaveRepository interface {
	Create(ctx context.Context, leave *LeaveRequest) error
	GetByID(ctx context.Context, id uuid.UUID) (*LeaveRequest, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status LeaveStatus, reviewerID uuid.UUID) error
	ListByEmployee(ctx context.Context, employeeID uuid.UUID) ([]*LeaveRequest, error)
	ListPending(ctx context.Context) ([]*LeaveRequest, error)
}

type GeofenceRepository interface {
	GetActiveZone(ctx context.Context) (*GeofenceZone, error)
}
