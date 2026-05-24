package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yourname/hr-portal-backend/internal/domain"
	myjwt "github.com/yourname/hr-portal-backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	empRepo   domain.EmployeeRepository
	jwtSecret string
}

func NewAuthUsecase(empRepo domain.EmployeeRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{empRepo: empRepo, jwtSecret: jwtSecret}
}

type RegisterRequest struct {
	FullName     string
	Email        string
	Password     string
	DepartmentID uuid.UUID
}

type LoginResponse struct {
	Token    string           `json:"token"`
	Employee *domain.Employee `json:"employee"`
}

func (u *AuthUsecase) Register(ctx context.Context, req RegisterRequest) (*domain.Employee, error) {
	existing, _ := u.empRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("bu email allaqachon ro'yxatdan o'tgan")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	emp := &domain.Employee{
		ID:           uuid.New(),
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         domain.RoleEmployee,
		DepartmentID: req.DepartmentID,
		CreatedAt:    time.Now(),
	}

	if err := u.empRepo.Create(ctx, emp); err != nil {
		return nil, err
	}
	return emp, nil
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	emp, err := u.empRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("email yoki parol noto'g'ri")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(emp.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("email yoki parol noto'g'ri")
	}

	jwtSvc := myjwt.NewJWTService(u.jwtSecret)
	token, err := jwtSvc.GenerateToken(emp.ID, string(emp.Role))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token, Employee: emp}, nil
}

func (u *AuthUsecase) ListEmployees(ctx context.Context) ([]*domain.Employee, error) {
	return u.empRepo.List(ctx)
}

func (u *AuthUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	return u.empRepo.GetByID(ctx, id)
}
