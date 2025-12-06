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

### 📅 Events

#### GET /events
List all events.

**Authentication:** Required

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "title": "Town Hall",
    "description": "Monthly Town Hall Meeting",
    "start_time": "2025-12-05T10:00:00Z",
    "end_time": "2025-12-05T11:00:00Z",
    "is_active": true,
    "created_at": "2025-12-04T20:00:00Z"
  }
]
```

#### POST /events
Create a new event.

**Authentication:** Required
**Role:** Admin only

**Request Body:**
```json
{
  "title": "Town Hall",
  "description": "Monthly Town Hall Meeting",
  "start_time": "2025-12-05T10:00:00Z",
  "end_time": "2025-12-05T11:00:00Z"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "title": "Town Hall",
  "description": "Monthly Town Hall Meeting",
  "start_time": "2025-12-05T10:00:00Z",
  "end_time": "2025-12-05T11:00:00Z",
  "is_active": true,
  "created_at": "2025-12-04T20:00:00Z"
}
```

---

### ⏰ Attendance & QR

### Get Event Attendance (Admin)
`GET /api/v1/events/:id/attendance`

Returns the list of attendance records for a specific event.

**Auth required**: Yes (Admin only)

**Response**:
```json
[
  {
    "id": 1,
    "user_id": 101,
    "user": {
      "id": 101,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john@example.com"
    },
    "event_id": 1,
    "check_in": "2025-12-05T09:05:00Z",
    "status": "present",
    "location": "Meeting Room A",
    "notes": "",
    "qr_token": "uuid-token"
  }
]
```

### Mark Manual Attendance (Admin)
`POST /api/v1/events/:id/attendance/manual`

Manually marks attendance for a user in an event.

**Auth required**: Yes (Admin only)

**Body**:
```json
{
  "user_id": 101,
  "notes": "Forgot phone"
}
```

**Response**:
```json
{
  "id": 2,
  "user_id": 101,
  "event_id": 1,
  "check_in": "2025-12-05T09:10:00Z",
  "status": "present",
  "location": "Manual Entry",
  "notes": "Forgot phone"
}
```

## QR Codes
#### GET /qr/active
Get the currently active QR codes for a specific event.

**Authentication:** Required  
**Role:** Admin only

**Query Parameters:**
- `event_id` (required)

**Response (200 OK):**
```json
{
  "qr_token": "abc123xyz789def456",
  "event_id": 1,
  "expires_at": "2025-12-04T22:40:00Z",
  "created_at": "2025-12-04T22:30:00Z",
  "is_active": true
}
```

---

#### POST /qr/generate
Generate a new QR code for an event.

**Authentication:** Required  
**Role:** Admin only

**Request Body:**
```json
{
  "event_id": 1
}
```

**Response (201 Created):**
```json
{
  "qr_token": "def456uvw012ghi789",
  "event_id": 1,
  "expires_at": "2025-12-04T22:50:00Z",
  "created_at": "2025-12-04T22:40:00Z",
  "is_active": true
}
```


---

#### POST /qr/deactivate
Deactivate the current QR code for an event.

**Authentication:** Required  
**Role:** Admin only

**Request Body:**
```json
{
  "event_id": 1
}
```

**Response (200 OK):**
```json
{
  "message": "QR code deactivated",
  "event_id": 1
}
```

---


#### POST /attendance/mark
Mark attendance using scanned QR token.

**Authentication:** Required

**Request Body:**
```json
{
  "qr_token": "abc123xyz789def456",
  "location": "Auditorium",
  "notes": "Scanned at entrance"
}
```

**Response (201 Created):**
```json
{
  "id": 15,
  "user_id": 5,
  "event_id": 1,
  "check_in": "2025-12-04T22:35:00Z",
  "status": "present",
  "notes": "Scanned at entrance",
  "location": "Auditorium",
  "created_at": "2025-12-04T22:35:00Z"
}
```

**Errors:**
- `400` - Invalid or expired QR token
- `400` - User already marked attendance for this event
- `404` - QR token not found

---

#### GET /attendance/history
Get attendance history for current user.

**Authentication:** Required

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": 5,
      "user_id": 1,
      "event_id": 1,
      "event": {
        "title": "Town Hall",
        "start_time": "2025-12-05T10:00:00Z"
      },
      "check_in": "2025-12-04T09:30:00Z",
      "status": "present",
      "notes": "Día completo",
      "location": "Oficina Central"
    }
  ],
  "total": 25,
  "page": 1,
  "limit": 10
}
```

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
| GET /events | - | ✅ | ✅ | ✅ |
| POST /events | - | - | - | ✅ |
| GET /qr/active | - | - | - | ✅ |
| POST /qr/generate | - | - | - | ✅ |
| POST /attendance/mark | - | ✅ | ✅ | ✅ |
| GET /attendance/history | - | ✅ | ✅ | ✅ |

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

## 🔗 Test Account

For development and testing:

```
Email: admin@example.com
Password: admin123
Role: admin
```

---

**Last Updated:** 2025-12-05  
**API Version:** v1  
**Backend Repository:** https://github.com/devjuank/attendance-backend
