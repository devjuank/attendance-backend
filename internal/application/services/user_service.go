package services

import (
	"errors"

	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"github.com/juank/attendance-backend/internal/domain/services"
	"github.com/juank/attendance-backend/pkg/utils"
)

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) services.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) Create(req *services.CreateUserRequest) (*models.User, error) {
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        req.Email,
		Password:     hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         req.Role,
		DepartmentID: req.DepartmentID,
		IsActive:     true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) GetByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserServiceImpl) GetByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *UserServiceImpl) Update(id uint, req *services.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Role != nil {
		user.Role = *req.Role
	}
	if req.DepartmentID != nil {
		user.DepartmentID = req.DepartmentID
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) Delete(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserServiceImpl) GetAll(page, limit int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.userRepo.GetAll(page, limit)
}

func (s *UserServiceImpl) ChangePassword(userID uint, req *services.ChangePasswordRequest) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(req.OldPassword, user.Password) {
		return errors.New("invalid old password")
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.Update(user)
}
