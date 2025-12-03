package repositories

import "github.com/juank/attendance-backend/internal/domain/models"

type DepartmentRepository interface {
	Create(department *models.Department) error
	GetByID(id uint) (*models.Department, error)
	GetAll() ([]models.Department, error)
	Update(department *models.Department) error
	Delete(id uint) error
}
