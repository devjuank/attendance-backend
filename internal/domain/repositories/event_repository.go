package repositories

import "github.com/juank/attendance-backend/internal/domain/models"

type EventRepository interface {
	Create(event *models.Event) error
	GetByID(id uint) (*models.Event, error)
	Update(event *models.Event) error
	Delete(id uint) error
	GetAll() ([]models.Event, error)
}
