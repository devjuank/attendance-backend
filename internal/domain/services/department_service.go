package services

import "github.com/juank/attendance-backend/internal/domain/models"

type CreateDepartmentRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	ManagerID   *uint  `json:"manager_id"`
}

type UpdateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ManagerID   *uint  `json:"manager_id"`
}

type DepartmentService interface {
	Create(req *CreateDepartmentRequest) (*models.Department, error)
	GetByID(id uint) (*models.Department, error)
	GetAll() ([]models.Department, error)
	Update(id uint, req *UpdateDepartmentRequest) (*models.Department, error)
	Delete(id uint) error
}
