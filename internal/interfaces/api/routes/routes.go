package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juank/attendance-backend/config"
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/interfaces/api/handlers"
	"github.com/juank/attendance-backend/internal/interfaces/api/middleware"
)

type Router struct {
	cfg               *config.Config
	authHandler       *handlers.AuthHandler
	userHandler       *handlers.UserHandler
	deptHandler       *handlers.DepartmentHandler
	attendanceHandler *handlers.AttendanceHandler
}

func NewRouter(
	cfg *config.Config,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	deptHandler *handlers.DepartmentHandler,
	attendanceHandler *handlers.AttendanceHandler,
) *Router {
	return &Router{
		cfg:               cfg,
		authHandler:       authHandler,
		userHandler:       userHandler,
		deptHandler:       deptHandler,
		attendanceHandler: attendanceHandler,
	}
}

func (r *Router) Setup(engine *gin.Engine) {
	// Global Middleware
	engine.Use(middleware.CORSMiddleware(r.cfg))
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// Health Check
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "attendance-backend",
			"version": "1.0.0",
		})
	})

	// API v1 Group
	v1 := engine.Group("/api/v1")
	{
		// Auth Routes (Public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", r.authHandler.Register)
			auth.POST("/login", r.authHandler.Login)
			auth.POST("/refresh", r.authHandler.RefreshToken)
			auth.POST("/logout", r.authHandler.Logout)
		}

		// Protected Routes
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware(r.cfg))
		{
			// User Routes
			users := protected.Group("/users")
			{
				users.GET("/me", r.userHandler.GetMe)
				users.PUT("/me/password", r.userHandler.ChangePassword)

				// Admin only - using middleware directly instead of sub-group
				users.POST("", middleware.RoleMiddleware(string(models.RoleAdmin)), r.userHandler.Create)
				users.GET("", middleware.RoleMiddleware(string(models.RoleAdmin)), r.userHandler.GetAll)
				users.GET("/:id", middleware.RoleMiddleware(string(models.RoleAdmin)), r.userHandler.GetByID)
				users.PUT("/:id", middleware.RoleMiddleware(string(models.RoleAdmin)), r.userHandler.Update)
				users.DELETE("/:id", middleware.RoleMiddleware(string(models.RoleAdmin)), r.userHandler.Delete)
			}

			// Department Routes
			departments := protected.Group("/departments")
			{
				departments.GET("", r.deptHandler.GetAll)
				departments.GET("/:id", r.deptHandler.GetByID)

				// Admin only
				departments.POST("", middleware.RoleMiddleware(string(models.RoleAdmin)), r.deptHandler.Create)
				departments.PUT("/:id", middleware.RoleMiddleware(string(models.RoleAdmin)), r.deptHandler.Update)
				departments.DELETE("/:id", middleware.RoleMiddleware(string(models.RoleAdmin)), r.deptHandler.Delete)
			}

			// Attendance Routes
			attendance := protected.Group("/attendance")
			{
				attendance.POST("/check-in", r.attendanceHandler.CheckIn)
				attendance.POST("/check-out", r.attendanceHandler.CheckOut)
				attendance.GET("/today", r.attendanceHandler.GetToday)
				attendance.GET("/history", r.attendanceHandler.GetMyHistory)
				attendance.GET("/range", r.attendanceHandler.GetByDateRange)
			}
		}
	}
}
