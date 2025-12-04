#!/bin/bash

# Script para probar el API del sistema de asistencia
# Asegúrate de que el servidor esté corriendo (make run)

BASE_URL="http://localhost:8080/api/v1"

echo "🧪 Probando API de Attendance System"
echo "===================================="
echo ""

# 1. Health Check
echo "1️⃣  Health Check"
curl -s http://localhost:8080/health | jq '.'
echo -e "\n"

# 2. Login con admin user (creado en migrations)
echo "2️⃣  Login (admin)"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }')

echo $LOGIN_RESPONSE | jq '.'
echo -e "\n"

# Extraer token
ACCESS_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.access_token')

if [ "$ACCESS_TOKEN" = "null" ] || [ -z "$ACCESS_TOKEN" ]; then
  echo "❌ Error: No se pudo obtener el token. ¿El servidor está corriendo?"
  exit 1
fi

echo "✅ Token obtenido: ${ACCESS_TOKEN:0:50}..."
echo -e "\n"

# 3. Obtener perfil del usuario actual
echo "3️⃣  Obtener mi perfil (GET /users/me)"
curl -s -X GET "$BASE_URL/users/me" \
  -H "Authorization: Bearer $ACCESS_TOKEN" | jq '.'
echo -e "\n"

# 4. Listar todos los departamentos
echo "4️⃣  Listar departamentos (GET /departments)"
curl -s -X GET "$BASE_URL/departments" \
  -H "Authorization: Bearer $ACCESS_TOKEN" | jq '.'
echo -e "\n"

# 5. Crear un nuevo departamento
echo "5️⃣  Crear departamento (POST /departments)"
DEPT_RESPONSE=$(curl -s -X POST "$BASE_URL/departments" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Tecnología",
    "description": "Departamento de TI"
  }')

echo $DEPT_RESPONSE | jq '.'
DEPT_ID=$(echo $DEPT_RESPONSE | jq -r '.ID')
echo "✅ Departamento creado con ID: $DEPT_ID"
echo -e "\n"

# 6. Registrar un nuevo usuario
echo "6️⃣  Registrar nuevo usuario (POST /auth/register)"
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "empleado@example.com",
    "password": "password123",
    "first_name": "Juan",
    "last_name": "Pérez"
  }')

echo $REGISTER_RESPONSE | jq '.'
echo -e "\n"

# 7. Check-in de asistencia
echo "7️⃣  Check-in (POST /attendance/check-in)"
CHECKIN_RESPONSE=$(curl -s -X POST "$BASE_URL/attendance/check-in" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "location": "Oficina Principal",
    "notes": "Llegada del día"
  }')

echo $CHECKIN_RESPONSE | jq '.'
echo -e "\n"

# 8. Ver asistencia de hoy
echo "8️⃣  Ver asistencia de hoy (GET /attendance/today)"
curl -s -X GET "$BASE_URL/attendance/today" \
  -H "Authorization: Bearer $ACCESS_TOKEN" | jq '.'
echo -e "\n"

# 9. Listar todos los usuarios (solo admin)
echo "9️⃣  Listar usuarios (GET /users - Admin only)"
curl -s -X GET "$BASE_URL/users?page=1&limit=10" \
  -H "Authorization: Bearer $ACCESS_TOKEN" | jq '.'
echo -e "\n"

echo "✅ Pruebas completadas!"
echo ""
echo "💡 Consejos:"
echo "   - Puedes modificar este script para probar otros endpoints"
echo "   - Usa 'jq' para formatear el JSON (brew install jq en Mac)"
echo "   - El token expira según JWT_EXPIRATION en .env"
