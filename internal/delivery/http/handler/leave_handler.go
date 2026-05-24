package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourname/hr-portal-backend/internal/domain"
	"github.com/yourname/hr-portal-backend/internal/usecase"
)

type LeaveHandler struct {
	uc *usecase.LeaveUsecase
}

func NewLeaveHandler(uc *usecase.LeaveUsecase) *LeaveHandler {
	return &LeaveHandler{uc: uc}
}

type applyLeaveBody struct {
	LeaveType string `json:"leave_type" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Reason    string `json:"reason" binding:"required"`
}

type reviewBody struct {
	Approved bool `json:"approved"`
}

func (h *LeaveHandler) Apply(c *gin.Context) {
	var body applyLeaveBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start, err := time.Parse("2006-01-02", body.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date formati: 2006-01-02"})
		return
	}
	end, err := time.Parse("2006-01-02", body.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_date formati: 2006-01-02"})
		return
	}

	empID, _ := c.Get("employee_id")
	leave, err := h.uc.Apply(c.Request.Context(), usecase.CreateLeaveRequest{
		EmployeeID: empID.(uuid.UUID),
		LeaveType:  domain.LeaveType(body.LeaveType),
		StartDate:  start,
		EndDate:    end,
		Reason:     body.Reason,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": leave})
}

func (h *LeaveHandler) MyLeaves(c *gin.Context) {
	empID, _ := c.Get("employee_id")
	leaves, err := h.uc.MyLeaves(c.Request.Context(), empID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": leaves})
}

func (h *LeaveHandler) ListPending(c *gin.Context) {
	leaves, err := h.uc.ListPending(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": leaves})
}

func (h *LeaveHandler) Review(c *gin.Context) {
	leaveID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id noto'g'ri"})
		return
	}

	var body reviewBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviewerID, _ := c.Get("employee_id")
	if err := h.uc.Review(c.Request.Context(), leaveID, reviewerID.(uuid.UUID), body.Approved); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "muvaffaqiyatli ko'rib chiqildi"})
}
