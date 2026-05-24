package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yourname/hr-portal-backend/internal/domain"
	"github.com/yourname/hr-portal-backend/internal/geofence"
)

type AttendanceUsecase struct {
	attendanceRepo domain.AttendanceRepository
	geofenceRepo   domain.GeofenceRepository
}

func NewAttendanceUsecase(ar domain.AttendanceRepository, gr domain.GeofenceRepository) *AttendanceUsecase {
	return &AttendanceUsecase{attendanceRepo: ar, geofenceRepo: gr}
}

type CheckInRequest struct {
	EmployeeID uuid.UUID
	Latitude   float64
	Longitude  float64
}

func (u *AttendanceUsecase) CheckIn(ctx context.Context, req CheckInRequest) (*domain.Attendance, error) {
	existing, err := u.attendanceRepo.GetTodayByEmployee(ctx, req.EmployeeID)
	if err == nil && existing != nil {
		return nil, errors.New("bugun allaqachon check-in qilingan")
	}

	zone, err := u.geofenceRepo.GetActiveZone(ctx)
	if err != nil {
		return nil, errors.New("geofence zone topilmadi")
	}

	if !geofence.IsInsideZone(req.Latitude, req.Longitude, zone.Latitude, zone.Longitude, zone.RadiusM) {
		return nil, errors.New("siz ish zonasidan tashqaridasiz")
	}

	now := time.Now()
	att := &domain.Attendance{
		ID:         uuid.New(),
		EmployeeID: req.EmployeeID,
		CheckIn:    &now,
		Status:     domain.StatusPresent,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		Date:       time.Now(),
	}

	if err := u.attendanceRepo.CheckIn(ctx, att); err != nil {
		return nil, err
	}
	return att, nil
}

func (u *AttendanceUsecase) CheckOut(ctx context.Context, employeeID uuid.UUID) error {
	_, err := u.attendanceRepo.GetTodayByEmployee(ctx, employeeID)
	if err != nil {
		return errors.New("bugun check-in topilmadi")
	}
	return u.attendanceRepo.CheckOut(ctx, employeeID)
}

func (u *AttendanceUsecase) GetToday(ctx context.Context, empID uuid.UUID) (*domain.Attendance, error) {
	return u.attendanceRepo.GetTodayByEmployee(ctx, empID)
}

func (u *AttendanceUsecase) MyAttendance(ctx context.Context, empID uuid.UUID) ([]*domain.Attendance, error) {
	return u.attendanceRepo.ListByEmployee(ctx, empID)
}

func (u *AttendanceUsecase) ListAll(ctx context.Context) ([]*domain.Attendance, error) {
	return u.attendanceRepo.ListAll(ctx)
}

func (u *AttendanceUsecase) GetByPeriod(ctx context.Context, empID uuid.UUID, period string) ([]*domain.Attendance, error) {
	now := time.Now()
	var from time.Time
	switch period {
	case "weekly":
		from = now.AddDate(0, 0, -7)
	case "monthly":
		from = now.AddDate(0, -1, 0)
	default:
		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	}
	return u.attendanceRepo.ListByEmployeePeriod(ctx, empID, from, now)
}
