package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juank/attendance-backend/config"
	"github.com/juank/attendance-backend/internal/application/services"
	"github.com/juank/attendance-backend/internal/infrastructure/database"
	"github.com/juank/attendance-backend/internal/infrastructure/persistence"
	"github.com/juank/attendance-backend/internal/interfaces/api/handlers"
	"github.com/juank/attendance-backend/internal/interfaces/api/routes"
	"github.com/juank/attendance-backend/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Cargar variables de entorno desde .env (si existe)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Cargar configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Inicializar logger
	if err := logger.InitLogger(cfg.Server.Env, cfg.Logging.Level); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Conectar a la base de datos
	db, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	logger.Info("Database connection established")

	logger.Info("Starting Attendance System API",
		zap.String("env", cfg.Server.Env),
		zap.String("port", cfg.Server.Port),
	)

	// Inicializar Repositorios
	userRepo := persistence.NewUserRepository(db)
	deptRepo := persistence.NewDepartmentRepository(db)
	attendanceRepo := persistence.NewAttendanceRepository(db)
	refreshTokenRepo := persistence.NewRefreshTokenRepository(db)

	// Inicializar Servicios
	authService := services.NewAuthService(userRepo, refreshTokenRepo, cfg)
	userService := services.NewUserService(userRepo)
	deptService := services.NewDepartmentService(deptRepo)
	attendanceService := services.NewAttendanceService(attendanceRepo)

	// Inicializar Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	deptHandler := handlers.NewDepartmentHandler(deptService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)

	// Configurar Gin según el entorno
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Crear router
	engine := gin.Default()

	// Configurar rutas
	router := routes.NewRouter(cfg, authHandler, userHandler, deptHandler, attendanceHandler)
	router.Setup(engine)

	// Configurar servidor
	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	logger.Info("Server is ready to handle requests",
		zap.String("address", serverAddr),
	)

	// Graceful shutdown
	go func() {
		if err := engine.Run(serverAddr); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Esperar señal de terminación
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	logger.Info("Server stopped")
}
