package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"uniqueIndex;not null;size:100" json:"name" validate:"required"`
	Description string         `gorm:"size:255" json:"description"`
	ManagerID   *uint          `json:"manager_id"`
	Manager     *User          `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	Users       []User         `gorm:"foreignKey:DepartmentID" json:"users,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
