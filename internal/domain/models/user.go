package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleManager  Role = "manager"
	RoleEmployee Role = "employee"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Email        string         `gorm:"uniqueIndex;not null;size:255" json:"email" validate:"required,email"`
	Password     string         `gorm:"not null" json:"-"`
	FirstName    string         `gorm:"not null;size:100" json:"first_name" validate:"required"`
	LastName     string         `gorm:"not null;size:100" json:"last_name" validate:"required"`
	Role         Role           `gorm:"type:varchar(20);not null;default:'employee'" json:"role"`
	DepartmentID *uint          `json:"department_id"`
	Department   *Department    `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
