package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" gorm:"unique" binding:"required,email"`
	Password  string    `json:"-" binding:"required"` // hash di service
	Role      string    `json:"role"`                 // e.g., "admin", "user"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
