package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID            uint              `gorm:"index;not null" json:"user_id"`
	TransactionNumber string            `gorm:"size:100;uniqueIndex;not null" json:"transaction_number"`
	TotalAmount       float64           `gorm:"not null" json:"total_amount"`
	PaymentMethod     string            `gorm:"size:100;not null;default:'Dompet Kampus'" json:"payment_method"`
	Status            string            `gorm:"size:50;default:'Selesai'" json:"status"`
	Items             []TransactionItem `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"items"`
}

type TransactionItem struct {
	gorm.Model
	TransactionID uint    `gorm:"index;not null" json:"transaction_id"`
	ProductID     uint    `gorm:"not null" json:"product_id"`
	ProductName   string  `gorm:"size:255;not null" json:"product_name"`
	Price         float64 `gorm:"not null" json:"price"`
	Quantity      int     `gorm:"not null" json:"quantity"`
	ImageUrl      string  `gorm:"size:512" json:"image_url"`
}

type CheckoutRequest struct {
	PaymentMethod string `json:"payment_method" binding:"required"`
}
