package repositories

import "github.com/juank/attendance-backend/internal/domain/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	GetAll(page, limit int) ([]models.User, int64, error)
}
