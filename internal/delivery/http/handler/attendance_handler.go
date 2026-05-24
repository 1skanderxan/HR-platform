package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourname/hr-portal-backend/internal/usecase"
)

type AttendanceHandler struct {
	uc *usecase.AttendanceUsecase
}

func NewAttendanceHandler(uc *usecase.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{uc: uc}
}

type checkInBody struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func (h *AttendanceHandler) CheckIn(c *gin.Context) {
	var body checkInBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	empID, _ := c.Get("employee_id")
	att, err := h.uc.CheckIn(c.Request.Context(), usecase.CheckInRequest{
		EmployeeID: empID.(uuid.UUID),
		Latitude:   body.Latitude,
		Longitude:  body.Longitude,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": att})
}

func (h *AttendanceHandler) CheckOut(c *gin.Context) {
	empID, _ := c.Get("employee_id")
	if err := h.uc.CheckOut(c.Request.Context(), empID.(uuid.UUID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "check-out muvaffaqiyatli"})
}

func (h *AttendanceHandler) GetToday(c *gin.Context) {
	empID, _ := c.Get("employee_id")
	att, err := h.uc.GetToday(c.Request.Context(), empID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": att})
}

func (h *AttendanceHandler) MyAttendance(c *gin.Context) {
	empID, _ := c.Get("employee_id")
	list, err := h.uc.MyAttendance(c.Request.Context(), empID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AttendanceHandler) ListAll(c *gin.Context) {
	list, err := h.uc.ListAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AttendanceHandler) GetByPeriod(c *gin.Context) {
	empID, _ := c.Get("employee_id")
	period := c.DefaultQuery("period", "daily")
	list, err := h.uc.GetByPeriod(c.Request.Context(), empID.(uuid.UUID), period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}
