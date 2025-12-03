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
	"github.com/juank/attendance-backend/internal/infrastructure/database"
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

	// Configurar Gin según el entorno
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Crear router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		dbStatus := "up"
		if err != nil || sqlDB.Ping() != nil {
			dbStatus = "down"
		}

		c.JSON(200, gin.H{
			"status":   "ok",
			"service":  "attendance-backend",
			"version":  "1.0.0",
			"database": dbStatus,
		})
	})

	// API v1 routes (placeholder)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// Configurar servidor
	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	logger.Info("Server is ready to handle requests",
		zap.String("address", serverAddr),
	)

	// Graceful shutdown
	go func() {
		if err := router.Run(serverAddr); err != nil {
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
