# Attendance System - Backend API

Sistema de gestión de asistencia desarrollado en Go con arquitectura limpia.

## 🚀 Características

- ✅ Autenticación JWT
- ✅ Gestión de usuarios y roles
- ✅ Registro de asistencia (check-in/check-out)
- ✅ Gestión de departamentos
- ✅ Generación de reportes
- ✅ API RESTful documentada con Swagger

## 🛠️ Stack Tecnológico

- **Go** 1.21+
- **Gin** - Framework web
- **GORM** - ORM
- **PostgreSQL** - Base de datos
- **JWT** - Autenticación
- **Viper** - Configuración
- **Zap** - Logging estructurado
- **Swagger** - Documentación API

## 📋 Requisitos Previos

- Go 1.21 o superior
- PostgreSQL 14+
- Make (opcional, para comandos útiles)

## 🔧 Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/juank/attendance-backend.git
cd attendance-backend
```

### 2. Configurar variables de entorno

```bash
cp .env.example .env
# Editar .env con tus configuraciones
```

### 3. Instalar dependencias

```bash
go mod download
```

### 4. Configurar base de datos

El proyecto utiliza Docker Compose para levantar una instancia local de PostgreSQL.

```bash
# Levantar PostgreSQL
make docker-up

# Verificar que la base de datos esté lista
docker ps
```

### 5. Ejecutar migraciones

Una vez que la base de datos esté corriendo:

```bash
# Ejecutar migraciones y seeds (crea usuario admin)
make migrate-up
```

### 6. Ejecutar el servidor

```bash
# Desarrollo
make run

# O con hot reload (si tienes Air instalado)
air
```

El servidor estará disponible en `http://localhost:8080`

## 📁 Estructura del Proyecto

```
attendance-backend/
├── cmd/
│   └── server/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── api/
│   │   ├── handlers/           # Controladores HTTP
│   │   ├── middleware/         # Middleware
│   │   └── routes/             # Definición de rutas
│   ├── domain/
│   │   ├── models/             # Modelos de dominio
│   │   └── repositories/       # Interfaces de repositorios
│   ├── infrastructure/
│   │   ├── database/           # Configuración de BD
│   │   └── persistence/        # Implementación de repositorios
│   ├── services/               # Lógica de negocio
│   └── utils/                  # Utilidades
├── config/                     # Archivos de configuración
├── migrations/                 # Migraciones de BD
├── docs/                       # Documentación Swagger
└── pkg/                        # Paquetes reutilizables
```

## 🔑 Variables de Entorno

Ver `.env.example` para todas las variables disponibles.

Variables principales:
- `PORT` - Puerto del servidor (default: 8080)
- `DB_HOST` - Host de PostgreSQL
- `DB_NAME` - Nombre de la base de datos
- `JWT_SECRET` - Secret para firmar tokens JWT
- `ALLOWED_ORIGINS` - Orígenes permitidos para CORS

## 📚 API Endpoints

### Autenticación
- `POST /api/v1/auth/register` - Registrar usuario
- `POST /api/v1/auth/login` - Iniciar sesión
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/logout` - Cerrar sesión

### Usuarios
- `GET /api/v1/users` - Listar usuarios (Admin)
- `GET /api/v1/users/:id` - Obtener usuario
- `PUT /api/v1/users/:id` - Actualizar usuario
- `DELETE /api/v1/users/:id` - Eliminar usuario (Admin)
- `GET /api/v1/users/me` - Perfil actual

### Asistencia
- `POST /api/v1/attendance/check-in` - Registrar entrada
- `POST /api/v1/attendance/check-out` - Registrar salida
- `GET /api/v1/attendance/me` - Mi historial
- `GET /api/v1/attendance/today` - Asistencia del día

### Departamentos
- `GET /api/v1/departments` - Listar departamentos
- `POST /api/v1/departments` - Crear departamento (Admin)
- `GET /api/v1/departments/:id` - Obtener departamento
- `PUT /api/v1/departments/:id` - Actualizar departamento
- `DELETE /api/v1/departments/:id` - Eliminar departamento

### Reportes
- `GET /api/v1/reports/attendance` - Reporte de asistencia
- `GET /api/v1/reports/user/:id` - Reporte por usuario
- `GET /api/v1/reports/department/:id` - Reporte por departamento

Ver documentación completa en `/swagger/index.html` cuando el servidor esté corriendo.

## 🧪 Testing

```bash
# Ejecutar todos los tests
make test

# Tests con coverage
make test-coverage

# Ver reporte de coverage
go tool cover -html=coverage.out
```

## 🐳 Docker

```bash
# Construir imagen
make docker-build

# Levantar servicios (app + PostgreSQL)
make docker-up

# Detener servicios
make docker-down
```

## 📖 Comandos Make Disponibles

```bash
make run              # Ejecutar servidor en desarrollo
make build            # Compilar binario
make test             # Ejecutar tests
make test-coverage    # Tests con coverage
make migrate-up       # Ejecutar migraciones
make migrate-down     # Revertir migraciones
make swagger          # Generar documentación Swagger
make lint             # Ejecutar linter
make docker-build     # Construir imagen Docker
make docker-up        # Levantar Docker Compose
make docker-down      # Detener Docker Compose
```

## 🔒 Seguridad

- Passwords hasheados con bcrypt
- Autenticación JWT con refresh tokens
- Rate limiting por IP
- Validación de inputs
- CORS configurado
- SQL injection prevention (GORM)

## 🚀 Deployment

### AWS (Recomendado)
1. Base de datos: AWS RDS (PostgreSQL)
2. Aplicación: AWS ECS o EC2
3. Secrets: AWS Secrets Manager
4. Load Balancer: AWS ALB

Ver `docs/deployment.md` para instrucciones detalladas.

## 📝 Documentación Adicional

- [ACTION_PLAN.md](ACTION_PLAN.md) - Plan de desarrollo
- [PROJECT_DESCRIPTION.md](PROJECT_DESCRIPTION.md) - Descripción técnica
- [API_CONTRACT.md](docs/API_CONTRACT.md) - Contrato de API

## 🤝 Contribuir

1. Fork el proyecto
2. Crear rama feature (`git checkout -b feature/AmazingFeature`)
3. Commit cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir Pull Request

## 📄 Licencia

Este proyecto es privado.

## 👥 Autores

- Juan K - Desarrollo inicial

## 🙏 Agradecimientos

- Frontend team por la integración
- Equipo de QA por el testing
