package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yourname/hr-portal-backend/internal/domain"
)

type attendanceRepo struct {
	db *sqlx.DB
}

func NewAttendanceRepository(db *sqlx.DB) domain.AttendanceRepository {
	return &attendanceRepo{db: db}
}

func (r *attendanceRepo) CheckIn(ctx context.Context, att *domain.Attendance) error {
	query := `INSERT INTO attendance (id, employee_id, check_in, status, latitude, longitude, date)
	          VALUES (:id, :employee_id, :check_in, :status, :latitude, :longitude, :date)`
	_, err := r.db.NamedExecContext(ctx, query, att)
	return err
}

func (r *attendanceRepo) CheckOut(ctx context.Context, employeeID uuid.UUID) error {
	now := time.Now()
	_, err := r.db.ExecContext(ctx,
		`UPDATE attendance SET check_out=$1
		 WHERE employee_id=$2 AND date=CURRENT_DATE AND check_out IS NULL`,
		now, employeeID)
	return err
}

func (r *attendanceRepo) GetTodayByEmployee(ctx context.Context, employeeID uuid.UUID) (*domain.Attendance, error) {
	att := &domain.Attendance{}
	err := r.db.GetContext(ctx, att,
		`SELECT a.*, COALESCE(e.full_name, '') as full_name
		 FROM attendance a
		 LEFT JOIN employees e ON a.employee_id = e.id
		 WHERE a.employee_id=$1 AND a.date=CURRENT_DATE`, employeeID)
	return att, err
}

func (r *attendanceRepo) ListByEmployee(ctx context.Context, employeeID uuid.UUID) ([]*domain.Attendance, error) {
	var list []*domain.Attendance
	err := r.db.SelectContext(ctx, &list,
		`SELECT a.*, COALESCE(e.full_name, '') as full_name
		 FROM attendance a
		 LEFT JOIN employees e ON a.employee_id = e.id
		 WHERE a.employee_id=$1 ORDER BY a.date DESC`, employeeID)
	return list, err
}

func (r *attendanceRepo) ListAll(ctx context.Context) ([]*domain.Attendance, error) {
	var list []*domain.Attendance
	err := r.db.SelectContext(ctx, &list,
		`SELECT a.*, COALESCE(e.full_name, '') as full_name
		 FROM attendance a
		 LEFT JOIN employees e ON a.employee_id = e.id
		 ORDER BY a.date DESC, a.check_in DESC`)
	return list, err
}

func (r *attendanceRepo) ListByEmployeePeriod(ctx context.Context, employeeID uuid.UUID, from, to time.Time) ([]*domain.Attendance, error) {
	var list []*domain.Attendance
	err := r.db.SelectContext(ctx, &list,
		`SELECT a.*, COALESCE(e.full_name, '') as full_name
		 FROM attendance a
		 LEFT JOIN employees e ON a.employee_id = e.id
		 WHERE a.employee_id=$1 AND a.date >= $2 AND a.date <= $3
		 ORDER BY a.date DESC`, employeeID, from, to)
	return list, err
}
