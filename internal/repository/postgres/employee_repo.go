package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yourname/hr-portal-backend/internal/domain"
)

type employeeRepo struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) domain.EmployeeRepository {
	return &employeeRepo{db: db}
}

func (r *employeeRepo) Create(ctx context.Context, emp *domain.Employee) error {
	query := `INSERT INTO employees (id, full_name, email, password_hash, role, department_id, created_at)
	          VALUES (:id, :full_name, :email, :password_hash, :role, :department_id, :created_at)`
	_, err := r.db.NamedExecContext(ctx, query, emp)
	return err
}

func (r *employeeRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	emp := &domain.Employee{}
	err := r.db.GetContext(ctx, emp, `SELECT * FROM employees WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (r *employeeRepo) GetByEmail(ctx context.Context, email string) (*domain.Employee, error) {
	emp := &domain.Employee{}
	err := r.db.GetContext(ctx, emp, `SELECT * FROM employees WHERE email=$1`, email)
	if err != nil {
		return nil, errors.New("topilmadi")
	}
	return emp, nil
}

func (r *employeeRepo) Update(ctx context.Context, emp *domain.Employee) error {
	_, err := r.db.NamedExecContext(ctx,
		`UPDATE employees SET full_name=:full_name, role=:role WHERE id=:id`, emp)
	return err
}

func (r *employeeRepo) List(ctx context.Context) ([]*domain.Employee, error) {
	var emps []*domain.Employee
	err := r.db.SelectContext(ctx, &emps, `SELECT * FROM employees ORDER BY created_at DESC`)
	return emps, err
}
