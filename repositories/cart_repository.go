package repositories

import (
	"github.com/SadamIRM/be-smoke-store/config"
	"github.com/SadamIRM/be-smoke-store/models"
)

type CartRepository struct{}

func NewCartRepository() *CartRepository {
	return &CartRepository{}
}

func (r *CartRepository) FindAllByUserID(userID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	result := config.DB.Where("user_id = ?", userID).Find(&items)
	return items, result.Error
}

func (r *CartRepository) FindByUserIDAndProductID(userID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	result := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (r *CartRepository) Create(item *models.CartItem) error {
	return config.DB.Create(item).Error
}

func (r *CartRepository) Update(item *models.CartItem) error {
	return config.DB.Save(item).Error
}

func (r *CartRepository) DeleteByProductID(userID, productID uint) error {
	return config.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.CartItem{}).Error
}

func (r *CartRepository) DeleteAllByUserID(userID uint) error {
	return config.DB.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error
}
