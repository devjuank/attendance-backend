package persistence

import (
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Department").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepositoryImpl) GetAll(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Preload("Department").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
