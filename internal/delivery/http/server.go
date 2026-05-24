package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/hr-portal-backend/config"
	"github.com/yourname/hr-portal-backend/internal/delivery/http/handler"
	"github.com/yourname/hr-portal-backend/internal/delivery/http/middleware"
	pg "github.com/yourname/hr-portal-backend/internal/repository/postgres"
	"github.com/yourname/hr-portal-backend/internal/usecase"
	myjwt "github.com/yourname/hr-portal-backend/pkg/jwt"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	db, err := pg.NewDB(cfg.DBUrl)
	if err != nil {
		panic("DB ga ulanib bo'lmadi: " + err.Error())
	}

	empRepo := pg.NewEmployeeRepository(db)
	attRepo := pg.NewAttendanceRepository(db)
	geoRepo := pg.NewGeofenceRepository(db)
	leaveRepo := pg.NewLeaveRepository(db)

	attUC := usecase.NewAttendanceUsecase(attRepo, geoRepo)
	leaveUC := usecase.NewLeaveUsecase(leaveRepo)
	authUC := usecase.NewAuthUsecase(empRepo, cfg.JWTSecret)

	jwtSvc := myjwt.NewJWTService(cfg.JWTSecret)

	attHandler := handler.NewAttendanceHandler(attUC)
	leaveHandler := handler.NewLeaveHandler(leaveUC)
	authHandler := handler.NewAuthHandler(authUC)

	r := gin.Default()

	api := r.Group("/api/v1")
	api.POST("/auth/login", authHandler.Login)
	api.POST("/auth/register", authHandler.Register)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtSvc))
	{
		protected.GET("/profile", authHandler.GetProfile)

		protected.POST("/attendance/checkin", attHandler.CheckIn)
		protected.POST("/attendance/checkout", attHandler.CheckOut)
		protected.GET("/attendance/today", attHandler.GetToday)
		protected.GET("/attendance/my", attHandler.MyAttendance)
		protected.GET("/attendance/period", attHandler.GetByPeriod)

		protected.POST("/leaves", leaveHandler.Apply)
		protected.GET("/leaves/my", leaveHandler.MyLeaves)

		hr := protected.Group("/")
		hr.Use(middleware.RequireRole("hr", "admin"))
		hr.GET("/attendance/all", attHandler.ListAll)
		hr.GET("/leaves/pending", leaveHandler.ListPending)
		hr.PUT("/leaves/:id/review", leaveHandler.Review)
		hr.GET("/employees", authHandler.ListEmployees)
		hr.POST("/employees", authHandler.Register)
	}

	return &Server{engine: r}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
