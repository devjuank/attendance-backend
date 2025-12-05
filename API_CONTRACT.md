# API Contract - Attendance System Backend

## Base URL
```
Development: http://localhost:8080/api/v1
Production: https://api.attendance-system.com/api/v1
```

## Authentication
Most endpoints require authentication using JWT Bearer tokens.

### Header Format
```
Authorization: Bearer <access_token>
```

---

## 📋 Endpoints

### 🔐 Authentication

#### POST /auth/register
Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "role": "employee",
  "department_id": null,
  "is_active": true,
  "created_at": "2025-12-04T20:00:00Z",
  "updated_at": "2025-12-04T20:00:00Z"
}
```

**Errors:**
- `400` - Email already registered or validation error

---

#### POST /auth/login
Login with email and password.

**Request Body:**
```json
{
  "email": "admin@example.com",
  "password": "admin123"
}
```

**Response (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Errors:**
- `401` - Invalid credentials
- `400` - Validation error

**Token Expiration:**
- Access Token: 24 hours (configurable)
- Refresh Token: 7 days (configurable)

---

#### POST /auth/refresh
Refresh access token using refresh token.

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Errors:**
- `401` - Invalid or expired refresh token

---

#### POST /auth/logout
Logout and invalidate refresh token.

**Authentication:** Required

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200 OK):**
```json
{
  "message": "logged out successfully"
}
```

---

### 👤 Users

#### GET /users/me
Get current user profile.

**Authentication:** Required

**Response (200 OK):**
```json
{
  "id": 1,
  "email": "admin@example.com",
  "first_name": "Admin",
  "last_name": "System",
  "role": "admin",
  "department_id": null,
  "is_active": true,
  "created_at": "2025-12-03T20:16:09Z",
  "updated_at": "2025-12-03T20:16:09Z"
}
```

---

#### PUT /users/me/password
Change own password.

**Authentication:** Required

**Request Body:**
```json
{
  "old_password": "currentPassword123",
  "new_password": "newPassword456"
}
```

**Response (200 OK):**
```json
{
  "message": "password changed successfully"
}
```

**Errors:**
- `400` - Invalid old password or validation error

---

#### GET /users
List all users (paginated).

**Authentication:** Required  
**Role:** Admin only

**Query Parameters:**
- `page` (optional, default: 1)
- `limit` (optional, default: 10, max: 100)

**Example:** `/users?page=1&limit=10`

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "email": "admin@example.com",
      "first_name": "Admin",
      "last_name": "System",
      "role": "admin",
      "department_id": null,
      "is_active": true,
      "created_at": "2025-12-03T20:16:09Z",
      "updated_at": "2025-12-03T20:16:09Z"
    }
  ],
  "total": 1,
  "page": 1,
  "limit": 10
}
```

**Errors:**
- `403` - Forbidden (non-admin user)

---

#### GET /users/:id
Get user by ID.

**Authentication:** Required  
**Role:** Admin only

**Response (200 OK):**
```json
{
  "id": 1,
  "email": "admin@example.com",
  "first_name": "Admin",
  "last_name": "System",
  "role": "admin",
  "department_id": null,
  "is_active": true,
  "created_at": "2025-12-03T20:16:09Z",
  "updated_at": "2025-12-03T20:16:09Z"
}
```

**Errors:**
- `404` - User not found
- `403` - Forbidden

---

#### POST /users
Create new user.

**Authentication:** Required  
**Role:** Admin only

**Request Body:**
```json
{
  "email": "newuser@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe",
  "role": "employee",
  "department_id": 1
}
```

**Response (201 Created):**
```json
{
  "id": 2,
  "email": "newuser@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "role": "employee",
  "department_id": 1,
  "is_active": true,
  "created_at": "2025-12-04T20:00:00Z",
  "updated_at": "2025-12-04T20:00:00Z"
}
```

**Roles:**
- `admin`
- `manager`
- `employee`

---

#### PUT /users/:id
Update user.

**Authentication:** Required  
**Role:** Admin only

**Request Body (all fields optional):**
```json
{
  "first_name": "Jane",
  "last_name": "Smith",
  "role": "manager",
  "department_id": 2,
  "is_active": false
}
```

**Response (200 OK):**
```json
{
  "id": 2,
  "email": "user@example.com",
  "first_name": "Jane",
  "last_name": "Smith",
  "role": "manager",
  "department_id": 2,
  "is_active": false,
  "created_at": "2025-12-04T20:00:00Z",
  "updated_at": "2025-12-04T20:05:00Z"
}
```

---

#### DELETE /users/:id
Delete user.

**Authentication:** Required  
**Role:** Admin only

**Response (200 OK):**
```json
{
  "message": "user deleted"
}
```

---

### 🏢 Departments

#### GET /departments
List all departments.

**Authentication:** Required

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Tecnología",
    "description": "Departamento de TI",
    "manager_id": null,
    "created_at": "2025-12-04T20:00:00Z",
    "updated_at": "2025-12-04T20:00:00Z"
  }
]
```

---

#### GET /departments/:id
Get department by ID.

**Authentication:** Required

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Tecnología",
  "description": "Departamento de TI",
  "manager_id": 5,
  "manager": {
    "id": 5,
    "email": "manager@example.com",
    "first_name": "Manager",
    "last_name": "Name"
  },
  "users": [
    {
      "id": 10,
      "email": "dev@example.com",
      "first_name": "Developer",
      "last_name": "Name"
    }
  ],
  "created_at": "2025-12-04T20:00:00Z",
  "updated_at": "2025-12-04T20:00:00Z"
}
```

---

#### POST /departments
Create department.

**Authentication:** Required  
**Role:** Admin only

**Request Body:**
```json
{
  "name": "Marketing",
  "description": "Departamento de Marketing",
  "manager_id": 3
}
```

**Response (201 Created):**
```json
{
  "id": 2,
  "name": "Marketing",
  "description": "Departamento de Marketing",
  "manager_id": 3,
  "created_at": "2025-12-04T20:00:00Z",
  "updated_at": "2025-12-04T20:00:00Z"
}
```

---

#### PUT /departments/:id
Update department.

**Authentication:** Required  
**Role:** Admin only

**Request Body (all fields optional):**
```json
{
  "name": "Marketing y Ventas",
  "description": "Updated description",
  "manager_id": 5
}
```

**Response (200 OK):**
```json
{
  "id": 2,
  "name": "Marketing y Ventas",
  "description": "Updated description",
  "manager_id": 5,
  "created_at": "2025-12-04T20:00:00Z",
  "updated_at": "2025-12-04T20:05:00Z"
}
```

---

#### DELETE /departments/:id
Delete department.

**Authentication:** Required  
**Role:** Admin only

**Response (200 OK):**
```json
{
  "message": "department deleted"
}
```

---

### ⏰ Attendance

#### GET /qr/active
Get the currently active QR code for attendance.

**Authentication:** Required  
**Role:** Admin only

**Response (200 OK):**
```json
{
  "qr_token": "abc123xyz789def456",
  "expires_at": "2025-12-04T22:40:00Z",
  "created_at": "2025-12-04T22:30:00Z",
  "is_active": true
}
```

**Response (404 Not Found):**
```json
{
  "error": "no active QR code found"
}
```

**Notes:**
- QR codes auto-expire after 10 minutes
- Only one QR code can be active at a time
- Backend auto-generates new QR if none exists or current expired

---

#### POST /qr/generate
Force generation of a new QR code (invalidates any previous active QR).

**Authentication:** Required  
**Role:** Admin only

**Response (201 Created):**
```json
{
  "qr_token": "def456uvw012ghi789",
  "expires_at": "2025-12-04T22:50:00Z",
  "created_at": "2025-12-04T22:40:00Z",
  "is_active": true
}
```

**Notes:**
- Invalidates any previously active QR code
- New QR is valid for 10 minutes from creation

---

#### POST /attendance/mark
Mark attendance using scanned QR token.

**Authentication:** Required

**Request Body:**
```json
{
  "qr_token": "abc123xyz789def456",
  "location": "Oficina Central",
  "notes": "Marked via QR scan"
}
```

**Response (201 Created):**
```json
{
  "id": 15,
  "user_id": 5,
  "check_in": "2025-12-04T22:35:00Z",
  "check_out": null,
  "status": "present",
  "notes": "Marked via QR scan",
  "location": "Oficina Central",
  "created_at": "2025-12-04T22:35:00Z",
  "updated_at": "2025-12-04T22:35:00Z"
}
```

**Status Values:**
- `present` - On time (before 9:15 AM)
- `late` - After 9:15 AM
- `absent` - No check-in

**Errors:**
- `400` - Invalid or expired QR token
- `400` - User already marked attendance today
- `404` - QR token not found

---

#### GET /attendance/today
Get today's attendance for current user.

**Authentication:** Required

**Response (200 OK):**
```json
{
  "id": 1,
  "user_id": 1,
  "check_in": "2025-12-04T09:30:00Z",
  "check_out": "2025-12-04T18:00:00Z",
  "status": "late",
  "notes": "Entrada del día",
  "location": "Oficina Central",
  "created_at": "2025-12-04T09:30:00Z",
  "updated_at": "2025-12-04T18:00:00Z"
}
```

**Errors:**
- `404` - No attendance record for today

---

#### GET /attendance/history
Get attendance history for current user (paginated).

**Authentication:** Required

**Query Parameters:**
- `page` (optional, default: 1)
- `limit` (optional, default: 10)

**Example:** `/attendance/history?page=1&limit=10`

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": 5,
      "user_id": 1,
      "check_in": "2025-12-04T09:30:00Z",
      "check_out": "2025-12-04T18:00:00Z",
      "status": "late",
      "notes": "Día completo",
      "location": "Oficina Central",
      "created_at": "2025-12-04T09:30:00Z",
      "updated_at": "2025-12-04T18:00:00Z"
    }
  ],
  "total": 25,
  "page": 1,
  "limit": 10
}
```

---

#### GET /attendance/range
Get attendance records for a date range.

**Authentication:** Required

**Query Parameters (required):**
- `start_date` - Format: YYYY-MM-DD
- `end_date` - Format: YYYY-MM-DD

**Example:** `/attendance/range?start_date=2025-12-01&end_date=2025-12-31`

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "check_in": "2025-12-01T09:00:00Z",
    "check_out": "2025-12-01T18:00:00Z",
    "status": "present",
    "notes": "",
    "location": "Oficina Central",
    "created_at": "2025-12-01T09:00:00Z",
    "updated_at": "2025-12-01T18:00:00Z"
  }
]
```

**Errors:**
- `400` - Invalid date format or missing parameters

---

## 🔒 Authorization Matrix

| Endpoint | Public | Employee | Manager | Admin |
|----------|--------|----------|---------|-------|
| POST /auth/register | ✅ | ✅ | ✅ | ✅ |
| POST /auth/login | ✅ | ✅ | ✅ | ✅ |
| POST /auth/refresh | ✅ | ✅ | ✅ | ✅ |
| POST /auth/logout | - | ✅ | ✅ | ✅ |
| GET /users/me | - | ✅ | ✅ | ✅ |
| PUT /users/me/password | - | ✅ | ✅ | ✅ |
| GET /users | - | - | - | ✅ |
| POST /users | - | - | - | ✅ |
| GET /users/:id | - | - | - | ✅ |
| PUT /users/:id | - | - | - | ✅ |
| DELETE /users/:id | - | - | - | ✅ |
| GET /departments | - | ✅ | ✅ | ✅ |
| GET /departments/:id | - | ✅ | ✅ | ✅ |
| POST /departments | - | - | - | ✅ |
| PUT /departments/:id | - | - | - | ✅ |
| DELETE /departments/:id | - | - | - | ✅ |
| GET /qr/active | - | - | - | ✅ |
| POST /qr/generate | - | - | - | ✅ |
| POST /attendance/mark | - | ✅ | ✅ | ✅ |
| GET /attendance/today | - | ✅ | ✅ | ✅ |
| GET /attendance/history | - | ✅ | ✅ | ✅ |
| GET /attendance/range | - | ✅ | ✅ | ✅ |

---

## 🚨 Error Responses

All error responses follow this format:

```json
{
  "error": "Detailed error message"
}
```

### Common HTTP Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error, business logic error)
- `401` - Unauthorized (missing or invalid token)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found
- `500` - Internal Server Error

---

## 📝 Notes

1. **Date/Time Format**: All timestamps are in ISO 8601 format (UTC)
2. **Pagination**: Default page size is 10, maximum is 100
3. **Authentication**: Include `Authorization: Bearer <token>` header for protected endpoints
4. **CORS**: Configured to allow requests from frontend origins
5. **Content-Type**: All requests/responses use `application/json`

---

## 🔗 Test Account

For development and testing:

```
Email: admin@example.com
Password: admin123
Role: admin
```

---

**Last Updated:** 2025-12-04  
**API Version:** v1  
**Backend Repository:** https://github.com/devjuank/attendance-backend
