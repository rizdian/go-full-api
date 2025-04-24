package model

import (
	"time"
)

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" binding:"required"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint      `json:"product_id" binding:"required"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" binding:"required,gt=0"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
