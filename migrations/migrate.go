package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juank/attendance-backend/config"
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/infrastructure/database"
	"github.com/juank/attendance-backend/pkg/logger"
	"gorm.io/gorm"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Parsear argumentos
	action := flag.String("action", "up", "Action to perform: up, down (not implemented for AutoMigrate)")
	flag.Parse()

	// Si no se pasa flag, buscar en argumentos posicionales (para compatibilidad con Makefile)
	if len(os.Args) > 1 && os.Args[1] != "-action" {
		*action = os.Args[1]
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

	// Conectar a la base de datos
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	switch *action {
	case "up":
		runMigrations(db)
	case "down":
		log.Println("Down migrations are not supported with AutoMigrate. Use SQL migrations if you need rollback capabilities.")
	default:
		log.Fatalf("Unknown action: %s", *action)
	}
}

func runMigrations(db *gorm.DB) {
	log.Println("Running migrations...")

	// Paso 1: Crear tablas sin restricciones de clave foránea para evitar problemas de dependencia circular
	// (User depende de Department, Department depende de User)
	log.Println("Step 1: Creating tables without foreign keys...")

	// Guardar configuración original
	originalConfig := db.Config.DisableForeignKeyConstraintWhenMigrating

	// Deshabilitar FKs temporalmente
	db.Config.DisableForeignKeyConstraintWhenMigrating = true

	if err := db.AutoMigrate(
		&models.User{},
		&models.Department{},
		&models.Event{},
		&models.Attendance{},
		&models.RefreshToken{},
		&models.QRCode{},
	); err != nil {
		log.Fatalf("Failed to run migrations (step 1): %v", err)
	}

	// Restaurar configuración original (habilitar FKs)
	db.Config.DisableForeignKeyConstraintWhenMigrating = originalConfig

	// Paso 2: Aplicar restricciones de clave foránea
	log.Println("Step 2: Applying foreign key constraints...")
	if err := db.AutoMigrate(
		&models.User{},
		&models.Department{},
		&models.Event{},
		&models.Attendance{},
		&models.RefreshToken{},
		&models.QRCode{},
	); err != nil {
		log.Fatalf("Failed to run migrations (step 2): %v", err)
	}

	log.Println("Migrations completed successfully")

	// Seed initial data if needed
	seedData(db)
}

func seedData(db *gorm.DB) {
	// Verificar si existe el admin
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count == 0 {
		log.Println("Seeding initial data...")

		// Crear admin por defecto (password: admin123)
		// Hash generado con bcrypt para "admin123"
		hashedPassword := "$2a$10$OuAf1TXdHzjtsfjF9mXwwObPzodmsuGUmVXWby.gtLS2OpY5cqIdy"

		admin := models.User{
			Email:     "admin@example.com",
			Password:  hashedPassword,
			FirstName: "Admin",
			LastName:  "System",
			Role:      models.RoleAdmin,
			IsActive:  true,
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Printf("Failed to seed admin user: %v", err)
		} else {
			log.Println("Admin user created successfully")
			log.Println("  Email: admin@example.com")
			log.Println("  Password: admin123")
		}
	}
}
