package repositories

import (
	"github.com/SadamIRM/be-smoke-store/config"
	"github.com/SadamIRM/be-smoke-store/models"
)

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(tx *models.Transaction) error {
	return config.DB.Create(tx).Error
}

func (r *TransactionRepository) FindAllByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := config.DB.Preload("Items").Where("user_id = ?", userID).Order("created_at desc").Find(&transactions)
	return transactions, result.Error
}
