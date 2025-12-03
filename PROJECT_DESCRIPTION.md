# Attendance System - Backend

## Descripción General

Backend RESTful API desarrollado en Go para un sistema de gestión de asistencia. Proporciona servicios de autenticación, gestión de usuarios, registro de asistencia y generación de reportes.

## Tecnologías Principales

- **Lenguaje**: Go 1.21+
- **Framework Web**: Gin (HTTP router de alto rendimiento)
- **Base de Datos**: PostgreSQL
- **ORM**: GORM
- **Autenticación**: JWT (JSON Web Tokens)
- **Documentación API**: Swagger/OpenAPI
- **Validación**: go-playground/validator
- **Configuración**: Viper
- **Logging**: Zap

## Arquitectura

### Estructura del Proyecto

```
attendance-backend/
├── cmd/
│   └── server/
│       └── main.go                 # Punto de entrada de la aplicación
├── internal/
│   ├── api/
│   │   ├── handlers/              # Controladores HTTP
│   │   ├── middleware/            # Middleware (auth, logging, cors)
│   │   └── routes/                # Definición de rutas
│   ├── domain/
│   │   ├── models/                # Modelos de dominio
│   │   └── repositories/          # Interfaces de repositorios
│   ├── infrastructure/
│   │   ├── database/              # Configuración de BD
│   │   └── persistence/           # Implementación de repositorios
│   ├── services/                  # Lógica de negocio
│   └── utils/                     # Utilidades (JWT, hash, etc.)
├── pkg/                           # Paquetes reutilizables
├── config/                        # Archivos de configuración
├── migrations/                    # Migraciones de BD
├── docs/                          # Documentación Swagger
├── .env.example                   # Ejemplo de variables de entorno
├── go.mod
├── go.sum
├── Makefile                       # Comandos útiles
└── README.md

```

### Patrones de Diseño

- **Clean Architecture**: Separación de capas (handlers, services, repositories)
- **Repository Pattern**: Abstracción de acceso a datos
- **Dependency Injection**: Inyección de dependencias para mejor testabilidad
- **Middleware Pattern**: Para autenticación, logging, CORS, etc.

## Características Principales

### 1. Autenticación y Autorización
- Registro de usuarios con validación
- Login con generación de JWT
- Refresh tokens para renovación de sesión
- Middleware de autenticación
- Control de acceso basado en roles (RBAC)

### 2. Gestión de Usuarios
- CRUD completo de usuarios
- Roles: Admin, Manager, Employee
- Perfil de usuario
- Cambio de contraseña

### 3. Registro de Asistencia
- Check-in/Check-out
- Registro de ubicación (opcional)
- Historial de asistencia
- Estados: Present, Absent, Late, On Leave

### 4. Gestión de Departamentos
- CRUD de departamentos
- Asignación de usuarios a departamentos

### 5. Reportes
- Reporte de asistencia por usuario
- Reporte de asistencia por departamento
- Reporte de asistencia por rango de fechas
- Exportación a CSV/PDF

### 6. Configuración del Sistema
- Horarios de trabajo
- Políticas de asistencia
- Notificaciones

## Modelos de Datos

### User
```go
type User struct {
    ID           uint      `gorm:"primaryKey"`
    Email        string    `gorm:"unique;not null"`
    Password     string    `gorm:"not null"`
    FirstName    string    `gorm:"not null"`
    LastName     string    `gorm:"not null"`
    Role         string    `gorm:"not null;default:'employee'"`
    DepartmentID *uint
    Department   *Department
    IsActive     bool      `gorm:"default:true"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

### Attendance
```go
type Attendance struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    User      User
    CheckIn   time.Time `gorm:"not null"`
    CheckOut  *time.Time
    Status    string    `gorm:"not null"` // present, late, absent, on_leave
    Notes     string
    Location  string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Department
```go
type Department struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"unique;not null"`
    Description string
    ManagerID   *uint
    Manager     *User
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## API Endpoints

### Autenticación
- `POST /api/v1/auth/register` - Registro de usuario
- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/refresh` - Refresh token
- `POST /api/v1/auth/logout` - Logout

### Usuarios
- `GET /api/v1/users` - Listar usuarios (Admin)
- `GET /api/v1/users/:id` - Obtener usuario
- `PUT /api/v1/users/:id` - Actualizar usuario
- `DELETE /api/v1/users/:id` - Eliminar usuario (Admin)
- `GET /api/v1/users/me` - Perfil del usuario actual
- `PUT /api/v1/users/me/password` - Cambiar contraseña

### Asistencia
- `POST /api/v1/attendance/check-in` - Registrar entrada
- `POST /api/v1/attendance/check-out` - Registrar salida
- `GET /api/v1/attendance/me` - Mi historial de asistencia
- `GET /api/v1/attendance/user/:id` - Historial de usuario (Manager/Admin)
- `GET /api/v1/attendance/today` - Asistencia del día

### Departamentos
- `GET /api/v1/departments` - Listar departamentos
- `POST /api/v1/departments` - Crear departamento (Admin)
- `GET /api/v1/departments/:id` - Obtener departamento
- `PUT /api/v1/departments/:id` - Actualizar departamento (Admin)
- `DELETE /api/v1/departments/:id` - Eliminar departamento (Admin)

### Reportes
- `GET /api/v1/reports/attendance` - Reporte de asistencia
- `GET /api/v1/reports/user/:id` - Reporte por usuario
- `GET /api/v1/reports/department/:id` - Reporte por departamento
- `GET /api/v1/reports/export` - Exportar reporte

## Seguridad

- Passwords hasheados con bcrypt
- JWT con expiración configurable
- CORS configurado para el frontend
- Rate limiting para prevenir ataques
- Validación de entrada en todos los endpoints
- SQL injection prevention (GORM)
- Logs de auditoría para acciones críticas

## Variables de Entorno

```env
# Server
PORT=8080
ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=attendance_db
DB_SSLMODE=disable

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h
JWT_REFRESH_EXPIRATION=168h

# CORS
ALLOWED_ORIGINS=http://localhost:5173

# Logging
LOG_LEVEL=info
```

## Deployment

### Desarrollo Local
- **Docker Compose**: PostgreSQL 14 + Adminer
- **Backend**: Ejecución local con `make run` o `air` (hot reload)
- **Migraciones**: Automáticas con GORM (custom script)

### Producción (AWS)
- **Base de Datos**: AWS RDS PostgreSQL (Multi-AZ)
- **Backend**: AWS ECS Fargate (Serverless Containers)
- **Registry**: AWS ECR para imágenes Docker
- **Load Balancer**: AWS ALB
- **IaC**: Terraform (en repositorio separado)
- **CI/CD**: GitHub Actions

## Testing

- Unit tests para services y handlers
- Integration tests para API endpoints
- Test coverage mínimo: 80%
- Mocks para repositorios

## Monitoreo y Logging

- Structured logging con Zap
- Health check endpoint: `GET /health`
- Metrics endpoint: `GET /metrics` (Prometheus)
- Error tracking y alertas

## Roadmap Futuro

- [ ] Notificaciones push
- [ ] Integración con sistemas de RRHH
- [ ] Reconocimiento facial para check-in
- [ ] Geofencing para validar ubicación
- [ ] Dashboard de analytics en tiempo real
- [ ] API GraphQL como alternativa
- [ ] Microservicios (si escala)
