package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID      uint    `gorm:"index;not null" json:"user_id"`
	ProductID   uint    `gorm:"not null" json:"product_id"`
	ProductName string  `gorm:"size:255;not null" json:"product_name"`
	Price       float64 `gorm:"not null" json:"price"`
	Quantity    int     `gorm:"not null;default:1" json:"quantity"`
	ImageUrl    string  `gorm:"size:512" json:"image_url"`
}

type AddCartItemRequest struct {
	ProductID   uint    `json:"product_id" binding:"required"`
	ProductName string  `json:"product_name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
	ImageUrl    string  `json:"image_url"`
}

type UpdateCartQuantityRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}
