package services

import "github.com/juank/attendance-backend/internal/domain/models"

type CreateUserRequest struct {
	Email        string      `json:"email" validate:"required,email"`
	Password     string      `json:"password" validate:"required,min=6"`
	FirstName    string      `json:"first_name" validate:"required"`
	LastName     string      `json:"last_name" validate:"required"`
	Role         models.Role `json:"role" validate:"required,oneof=admin manager employee"`
	DepartmentID *uint       `json:"department_id"`
}

type UpdateUserRequest struct {
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Role         *models.Role `json:"role" validate:"omitempty,oneof=admin manager employee"`
	DepartmentID *uint        `json:"department_id"`
	IsActive     *bool        `json:"is_active"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type UserService interface {
	Create(req *CreateUserRequest) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(id uint, req *UpdateUserRequest) (*models.User, error)
	Delete(id uint) error
	GetAll(page, limit int) ([]models.User, int64, error)
	ChangePassword(userID uint, req *ChangePasswordRequest) error
}
