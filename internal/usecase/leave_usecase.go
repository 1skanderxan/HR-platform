package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yourname/hr-portal-backend/internal/domain"
)

type LeaveUsecase struct {
	leaveRepo domain.LeaveRepository
}

func NewLeaveUsecase(lr domain.LeaveRepository) *LeaveUsecase {
	return &LeaveUsecase{leaveRepo: lr}
}

type CreateLeaveRequest struct {
	EmployeeID uuid.UUID
	LeaveType  domain.LeaveType
	StartDate  time.Time
	EndDate    time.Time
	Reason     string
}

func (u *LeaveUsecase) Apply(ctx context.Context, req CreateLeaveRequest) (*domain.LeaveRequest, error) {
	if req.EndDate.Before(req.StartDate) {
		return nil, errors.New("tugash sanasi boshlash sanasidan oldin bo'lishi mumkin emas")
	}

	leave := &domain.LeaveRequest{
		ID:         uuid.New(),
		EmployeeID: req.EmployeeID,
		LeaveType:  req.LeaveType,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Reason:     req.Reason,
		Status:     domain.LeaveStatusPending,
		CreatedAt:  time.Now(),
	}

	if err := u.leaveRepo.Create(ctx, leave); err != nil {
		return nil, err
	}
	return leave, nil
}

func (u *LeaveUsecase) Review(ctx context.Context, leaveID, reviewerID uuid.UUID, approved bool) error {
	leave, err := u.leaveRepo.GetByID(ctx, leaveID)
	if err != nil {
		return errors.New("so'rov topilmadi")
	}
	if leave.Status != domain.LeaveStatusPending {
		return errors.New("bu so'rov allaqachon ko'rib chiqilgan")
	}

	status := domain.LeaveStatusRejected
	if approved {
		status = domain.LeaveStatusApproved
	}
	return u.leaveRepo.UpdateStatus(ctx, leaveID, status, reviewerID)
}

func (u *LeaveUsecase) MyLeaves(ctx context.Context, employeeID uuid.UUID) ([]*domain.LeaveRequest, error) {
	return u.leaveRepo.ListByEmployee(ctx, employeeID)
}

func (u *LeaveUsecase) ListPending(ctx context.Context) ([]*domain.LeaveRequest, error) {
	return u.leaveRepo.ListPending(ctx)
}
