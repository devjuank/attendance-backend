package persistence

import (
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type DepartmentRepositoryImpl struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) repositories.DepartmentRepository {
	return &DepartmentRepositoryImpl{db: db}
}

func (r *DepartmentRepositoryImpl) Create(department *models.Department) error {
	return r.db.Create(department).Error
}

func (r *DepartmentRepositoryImpl) GetByID(id uint) (*models.Department, error) {
	var department models.Department
	if err := r.db.Preload("Manager").Preload("Users").First(&department, id).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *DepartmentRepositoryImpl) GetAll() ([]models.Department, error) {
	var departments []models.Department
	if err := r.db.Preload("Manager").Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DepartmentRepositoryImpl) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *DepartmentRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}
