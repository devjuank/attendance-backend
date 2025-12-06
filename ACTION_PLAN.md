# Action Plan - Attendance System Backend (Go)

## Fase 1: Configuración Inicial del Proyecto ✅

### 1.1 Inicialización del Proyecto
- [ ] Inicializar módulo Go (`go mod init`)
- [ ] Crear estructura de directorios según Clean Architecture
- [ ] Configurar `.gitignore` para Go
- [ ] Crear `README.md` con instrucciones de setup
- [ ] Crear `.env.example` con variables de entorno necesarias

### 1.2 Instalación de Dependencias Core
- [ ] Instalar Gin framework (`github.com/gin-gonic/gin`)
- [ ] Instalar GORM (`gorm.io/gorm`)
- [ ] Instalar driver PostgreSQL (`gorm.io/driver/postgres`)
- [ ] Instalar Viper para configuración (`github.com/spf13/viper`)
- [ ] Instalar Zap para logging (`go.uber.org/zap`)
- [ ] Instalar JWT library (`github.com/golang-jwt/jwt/v5`)
- [ ] Instalar bcrypt (`golang.org/x/crypto/bcrypt`)
- [ ] Instalar validator (`github.com/go-playground/validator/v10`)
- [ ] Instalar godotenv (`github.com/joho/godotenv`)

### 1.3 Configuración Base
- [ ] Crear sistema de configuración con Viper
- [ ] Configurar logger con Zap
- [ ] Crear archivo de configuración para diferentes entornos (dev, prod)

---

## Fase 2: Configuración de Base de Datos

### 2.1 Setup de PostgreSQL
### 2.1 Setup de PostgreSQL
- [x] Crear script de inicialización de BD
- [x] Configurar conexión a PostgreSQL con GORM
- [x] Implementar pool de conexiones
- [x] Crear health check para BD

### 2.2 Modelos de Datos
- [x] Crear modelo `User` con GORM tags
- [x] Crear modelo `Department`
- [x] Crear modelo `Attendance`
- [x] Crear modelo `RefreshToken`
- [x] Definir relaciones entre modelos

### 2.3 Migraciones
- [x] Configurar sistema de migraciones
- [x] Crear migración inicial para tabla `users`
- [x] Crear migración para tabla `departments`
- [x] Crear migración para tabla `attendances`
- [x] Crear migración para tabla `refresh_tokens`
- [x] Crear seeds para datos iniciales (admin user)

---

## Fase 3: Capa de Infraestructura ✅

### 3.1 Repositorios - Interfaces
- [x] Definir interface `UserRepository`
- [x] Definir interface `DepartmentRepository`
- [x] Definir interface `AttendanceRepository`
- [x] Definir interface `RefreshTokenRepository`

### 3.2 Repositorios - Implementación
- [x] Implementar `UserRepositoryImpl` con GORM
- [x] Implementar `DepartmentRepositoryImpl` con GORM
- [x] Implementar `AttendanceRepositoryImpl` con GORM
- [x] Implementar `RefreshTokenRepositoryImpl` con GORM

---

## Fase 4: Capa de Servicios (Lógica de Negocio) ✅

### 4.1 Utilidades
- [x] Crear utilidad para hash de passwords (bcrypt)
- [x] Crear utilidad para generación de JWT
- [x] Crear utilidad para validación de JWT
- [x] Crear utilidad para refresh tokens
- [x] Crear validadores personalizados (via go-playground/validator)

### 4.2 Auth Service
- [x] Implementar `Register` (validación, hash, creación de usuario)
- [x] Implementar `Login` (validación, verificación, generación de tokens)
- [x] Implementar `RefreshToken` (validación y renovación)
- [x] Implementar `Logout` (invalidación de tokens)

### 4.3 User Service
- [x] Implementar `GetAllUsers` (con paginación)
- [x] Implementar `GetUserByID`
- [x] Implementar `GetUserByEmail`
- [x] Implementar `UpdateUser`
- [x] Implementar `DeleteUser`
- [x] Implementar `ChangePassword`
- [x] Implementar `GetProfile`

### 4.4 Department Service
- [x] Implementar CRUD completo de departamentos
- [x] Implementar asignación de usuarios a departamentos
- [x] Implementar listado de usuarios por departamento

### 4.5 Attendance Service
- [x] Implementar `CheckIn` (validaciones de negocio)
- [x] Implementar `CheckOut`
- [x] Implementar `GetAttendanceHistory` (con filtros)
- [x] Implementar `GetTodayAttendance`
- [x] Implementar `GetUserAttendance`
- [x] Implementar cálculo de estado (present, late, absent)

### 4.6 Report Service
- [ ] Implementar generación de reportes por usuario
- [ ] Implementar generación de reportes por departamento
- [ ] Implementar generación de reportes por rango de fechas
- [ ] Implementar exportación a CSV
- [ ] Implementar cálculo de estadísticas (horas trabajadas, etc.)

---

## Fase 5: Capa de API (Handlers y Routes) ✅

### 5.1 Middleware
- [x] Crear middleware de autenticación JWT
- [x] Crear middleware de autorización por roles
- [x] Crear middleware de logging de requests (Gin default)
- [x] Crear middleware de CORS
- [ ] Crear middleware de rate limiting
- [x] Crear middleware de recovery (panic handler - Gin default)
- [x] Crear middleware de validación de request body (Gin default)

### 5.2 DTOs (Data Transfer Objects)
- [x] Crear DTOs para Auth (RegisterRequest, LoginRequest, LoginResponse)
- [x] Crear DTOs para User (CreateUserRequest, UpdateUserRequest, UserResponse)
- [x] Crear DTOs para Attendance (CheckInRequest, CheckOutRequest, AttendanceResponse)
- [x] Crear DTOs para Department (CreateDepartmentRequest, etc.)
- [ ] Crear DTOs para Reports (ReportRequest, ReportResponse)

### 5.3 Handlers - Auth
- [x] Implementar `POST /api/v1/auth/register`
- [x] Implementar `POST /api/v1/auth/login`
- [x] Implementar `POST /api/v1/auth/refresh`
- [x] Implementar `POST /api/v1/auth/logout`

### 5.4 Handlers - Users
- [x] Implementar `GET /api/v1/users` (Admin only)
- [x] Implementar `GET /api/v1/users/:id`
- [x] Implementar `PUT /api/v1/users/:id`
- [x] Implementar `DELETE /api/v1/users/:id` (Admin only)
- [x] Implementar `GET /api/v1/users/me`
- [x] Implementar `PUT /api/v1/users/me/password`

### 5.5 Handlers - Attendance
- [x] Implementar `POST /api/v1/attendance/check-in`
- [x] Implementar `POST /api/v1/attendance/check-out`
- [x] Implementar `GET /api/v1/attendance/today`
- [x] Implementar `GET /api/v1/attendance/history`
- [x] Implementar `GET /api/v1/attendance/range` (por fecha)y`

### 5.6 Handlers - Departments
- [x] Implementar `GET /api/v1/departments` (listar todos)
- [x] Implementar `GET /api/v1/departments/:id`
- [x] Implementar `POST /api/v1/departments` (Admin only)
- [x] Implementar `PUT /api/v1/departments/:id` (Admin only)
- [x] Implementar `DELETE /api/v1/departments/:id` (Admin only)

### 5.7 Handlers - Reports
- [ ] Implementar `GET /api/v1/reports/attendance`
- [ ] Implementar `GET /api/v1/reports/user/:id`
- [ ] Implementar `GET /api/v1/reports/department/:id`
- [ ] Implementar `GET /api/v1/reports/export`

### 5.8 Routes Setup
- [x] Configurar router principal
- [x] Agrupar rutas por recurso
- [x] Aplicar middleware de autenticación a rutas protegidas
- [x] Aplicar middleware de autorización (roles)
- [x] Documentar endpoints en READMEhealth check (`GET /health`)

---

## Fase 6: Documentación de API

### 6.1 Swagger/OpenAPI
- [ ] Instalar Swag (`github.com/swaggo/swag`)
- [ ] Instalar gin-swagger (`github.com/swaggo/gin-swagger`)
- [ ] Anotar handlers con comentarios Swagger
- [ ] Generar documentación Swagger
- [ ] Configurar endpoint `/swagger/*` para UI de Swagger

### 6.2 Documentación General
- [ ] Crear `API_CONTRACT.md` con especificación completa
- [ ] Documentar ejemplos de requests/responses
- [ ] Documentar códigos de error
- [ ] Crear guía de autenticación

---

## Fase 7: Testing

### 7.1 Unit Tests
- [ ] Tests para utilidades (JWT, hash, validators)
- [ ] Tests para servicios de autenticación
- [ ] Tests para servicios de usuarios
- [ ] Tests para servicios de asistencia
- [ ] Tests para servicios de departamentos
- [ ] Configurar mocks para repositorios

### 7.2 Integration Tests
- [ ] Tests de endpoints de autenticación
- [ ] Tests de endpoints de usuarios
- [ ] Tests de endpoints de asistencia
- [ ] Tests de endpoints de departamentos
- [ ] Tests de middleware de autenticación
- [ ] Setup de BD de prueba

### 7.3 Coverage
- [ ] Configurar herramienta de coverage
- [ ] Alcanzar mínimo 80% de coverage
- [ ] Generar reportes de coverage

---

## Fase 8: Optimización y Seguridad

### 8.1 Seguridad
- [ ] Implementar rate limiting por IP
- [ ] Implementar validación estricta de inputs
- [ ] Configurar CORS correctamente
- [ ] Implementar headers de seguridad (helmet)
- [ ] Auditoría de dependencias
- [ ] Implementar logs de auditoría para acciones críticas

### 8.2 Performance
- [ ] Implementar paginación en listados
- [ ] Optimizar queries de BD (índices)
- [ ] Implementar caching (Redis) para tokens
- [ ] Optimizar serialización JSON
- [ ] Profiling de performance

### 8.3 Monitoring
- [ ] Implementar métricas con Prometheus
- [ ] Configurar health checks detallados
- [ ] Implementar structured logging
- [ ] Configurar alertas básicas

---

## Fase 9: Containerización y Deploy

### 9.1 Docker
- [ ] Crear `Dockerfile` optimizado (multi-stage build)
- [ ] Crear `docker-compose.yml` para desarrollo local
- [ ] Configurar hot-reload para desarrollo
- [ ] Crear script de inicialización de BD en Docker

### 9.2 CI/CD
- [ ] Configurar GitHub Actions para tests
- [ ] Configurar build automático
- [ ] Configurar linting (golangci-lint)
- [ ] Configurar security scanning

### 9.3 Deployment a AWS
- [ ] Configurar AWS RDS para PostgreSQL
- [ ] Configurar AWS Secrets Manager para secrets
- [ ] Deploy en AWS ECS o EC2
- [ ] Configurar Load Balancer
- [ ] Configurar dominio y SSL/TLS
- [ ] Configurar backups automáticos de BD

---

## Fase 10: Integración con Frontend

### 10.1 CORS y Configuración
- [ ] Configurar CORS para dominio del frontend
- [ ] Validar que todos los endpoints estén accesibles
- [ ] Configurar headers necesarios

### 10.2 Testing de Integración
- [ ] Probar flujo completo de autenticación
- [ ] Probar flujo de registro de asistencia
- [ ] Probar generación de reportes
- [ ] Validar manejo de errores
- [ ] Validar formato de respuestas

### 10.3 Documentación para Frontend
- [ ] Compartir documentación Swagger
- [ ] Documentar flujos de autenticación
- [ ] Documentar manejo de tokens
- [ ] Proveer ejemplos de integración

---

## Fase 11: Refinamiento y Lanzamiento

### 11.1 Code Review
- [ ] Revisar código siguiendo Go best practices
- [ ] Refactorizar código duplicado
- [ ] Mejorar nombres de variables y funciones
- [ ] Agregar comentarios donde sea necesario

### 11.2 Documentación Final
- [ ] Actualizar README.md con instrucciones completas
- [ ] Documentar variables de entorno
- [ ] Crear guía de deployment
- [ ] Crear guía de troubleshooting

### 11.3 Preparación para Producción
- [ ] Configurar logs de producción
- [ ] Configurar monitoreo
- [ ] Preparar plan de rollback
- [ ] Realizar pruebas de carga
- [ ] Validar backups

### 11.4 Lanzamiento
- [ ] Deploy a staging
- [ ] Testing en staging
- [ ] Deploy a producción
- [ ] Monitoreo post-deploy
- [ ] Documentar lecciones aprendidas

---

## Makefile - Comandos Útiles

```makefile
# Desarrollo
run:
	go run cmd/server/main.go

# Testing
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Database
migrate-up:
	go run migrations/migrate.go up

migrate-down:
	go run migrations/migrate.go down

# Docker
docker-build:
	docker build -t attendance-backend .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Swagger
swagger:
	swag init -g cmd/server/main.go

# Linting
lint:
	golangci-lint run

# Build
build:
	go build -o bin/server cmd/server/main.go
```

---

## Notas Importantes

- **Prioridad**: Seguir el orden de las fases para evitar dependencias faltantes
- **Testing**: Escribir tests a medida que se desarrolla cada feature
- **Commits**: Hacer commits frecuentes y descriptivos
- **Documentación**: Mantener documentación actualizada
- **Code Review**: Revisar código antes de merge a main
- **Seguridad**: Nunca commitear secrets o .env files

## Estimación de Tiempo

- **Fase 1-2**: 1-2 días
- **Fase 3-4**: 3-4 días
- **Fase 5**: 3-4 días
- **Fase 6-7**: 2-3 días
- **Fase 8-9**: 2-3 días
- **Fase 10-11**: 2-3 días


## Fase 12: Arquitectura Basada en Eventos (Nuevo) ✅

### 12.1 Modelos y Base de Datos
- [x] Crear modelo `Event`
- [x] Actualizar modelo `QRCode` con `EventID`
- [x] Actualizar modelo `Attendance` con `EventID`
- [x] Crear repositorio `EventRepository`
- [x] Actualizar repositorios existentes

### 12.2 Lógica de Negocio
- [x] Crear `EventService`
- [x] Actualizar `QRService` para soportar eventos
- [x] Actualizar `AttendanceService` para validar evento

### 12.3 API
- [x] Crear `EventHandler`
- [x] Actualizar `QRHandler`
- [x] Registrar rutas de eventos
- [x] Actualizar documentación (`API_CONTRACT.md`, `PROJECT_DESCRIPTION.md`)
