package services

import (
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
)

type EventService struct {
	eventRepo repositories.EventRepository
}

func NewEventService(eventRepo repositories.EventRepository) *EventService {
	return &EventService{eventRepo: eventRepo}
}

func (s *EventService) Create(event *models.Event) error {
	return s.eventRepo.Create(event)
}

func (s *EventService) GetByID(id uint) (*models.Event, error) {
	return s.eventRepo.GetByID(id)
}

func (s *EventService) Update(event *models.Event) error {
	return s.eventRepo.Update(event)
}

func (s *EventService) Delete(id uint) error {
	return s.eventRepo.Delete(id)
}

func (s *EventService) GetAll() ([]models.Event, error) {
	return s.eventRepo.GetAll()
}
