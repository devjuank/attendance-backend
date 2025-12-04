package services

import (
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"github.com/juank/attendance-backend/internal/domain/services"
)

type DepartmentServiceImpl struct {
	deptRepo repositories.DepartmentRepository
}

func NewDepartmentService(deptRepo repositories.DepartmentRepository) services.DepartmentService {
	return &DepartmentServiceImpl{
		deptRepo: deptRepo,
	}
}

func (s *DepartmentServiceImpl) Create(req *services.CreateDepartmentRequest) (*models.Department, error) {
	dept := &models.Department{
		Name:        req.Name,
		Description: req.Description,
		ManagerID:   req.ManagerID,
	}

	if err := s.deptRepo.Create(dept); err != nil {
		return nil, err
	}

	return dept, nil
}

func (s *DepartmentServiceImpl) GetByID(id uint) (*models.Department, error) {
	return s.deptRepo.GetByID(id)
}

func (s *DepartmentServiceImpl) GetAll() ([]models.Department, error) {
	return s.deptRepo.GetAll()
}

func (s *DepartmentServiceImpl) Update(id uint, req *services.UpdateDepartmentRequest) (*models.Department, error) {
	dept, err := s.deptRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		dept.Name = req.Name
	}
	if req.Description != "" {
		dept.Description = req.Description
	}
	if req.ManagerID != nil {
		dept.ManagerID = req.ManagerID
	}

	if err := s.deptRepo.Update(dept); err != nil {
		return nil, err
	}

	return dept, nil
}

func (s *DepartmentServiceImpl) Delete(id uint) error {
	return s.deptRepo.Delete(id)
}
