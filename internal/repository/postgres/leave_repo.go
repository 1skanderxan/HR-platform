package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yourname/hr-portal-backend/internal/domain"
)

type leaveRepo struct {
	db *sqlx.DB
}

func NewLeaveRepository(db *sqlx.DB) domain.LeaveRepository {
	return &leaveRepo{db: db}
}

func (r *leaveRepo) Create(ctx context.Context, leave *domain.LeaveRequest) error {
	query := `INSERT INTO leave_requests 
              (id, employee_id, leave_type, start_date, end_date, reason, status, created_at)
              VALUES (:id, :employee_id, :leave_type, :start_date, :end_date, :reason, :status, :created_at)`
	_, err := r.db.NamedExecContext(ctx, query, leave)
	return err
}

func (r *leaveRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.LeaveRequest, error) {
	leave := &domain.LeaveRequest{}
	err := r.db.GetContext(ctx, leave,
		`SELECT l.*, COALESCE(e.full_name, '') as full_name
		 FROM leave_requests l
		 LEFT JOIN employees e ON l.employee_id = e.id
		 WHERE l.id = $1`, id)
	return leave, err
}

func (r *leaveRepo) UpdateStatus(ctx context.Context, id uuid.UUID, status domain.LeaveStatus, reviewerID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE leave_requests SET status = $1, reviewer_id = $2 WHERE id = $3`,
		status, reviewerID, id)
	return err
}

func (r *leaveRepo) ListByEmployee(ctx context.Context, employeeID uuid.UUID) ([]*domain.LeaveRequest, error) {
	var list []*domain.LeaveRequest
	err := r.db.SelectContext(ctx, &list,
		`SELECT l.*, COALESCE(e.full_name, '') as full_name
		 FROM leave_requests l
		 LEFT JOIN employees e ON l.employee_id = e.id
		 WHERE l.employee_id = $1 
		 ORDER BY l.created_at DESC`, employeeID)
	return list, err
}

func (r *leaveRepo) ListPending(ctx context.Context) ([]*domain.LeaveRequest, error) {
	var list []*domain.LeaveRequest
	err := r.db.SelectContext(ctx, &list,
		`SELECT l.*, COALESCE(e.full_name, '') as full_name
		 FROM leave_requests l
		 LEFT JOIN employees e ON l.employee_id = e.id
		 WHERE l.status = 'pending' 
		 ORDER BY l.created_at ASC`)
	return list, err
}
