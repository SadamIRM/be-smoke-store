package services

import (
	"errors"

	"github.com/SadamIRM/be-smoke-store/models"
	"github.com/SadamIRM/be-smoke-store/repositories"
	"gorm.io/gorm"
)

type CartService struct {
	cartRepo *repositories.CartRepository
}

func NewCartService() *CartService {
	return &CartService{cartRepo: repositories.NewCartRepository()}
}

func (s *CartService) GetCart(userID uint) ([]models.CartItem, error) {
	return s.cartRepo.FindAllByUserID(userID)
}

func (s *CartService) AddToCart(userID uint, req *models.AddCartItemRequest) (*models.CartItem, error) {
	item, err := s.cartRepo.FindByUserIDAndProductID(userID, req.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Item baru
			newItem := &models.CartItem{
				UserID:      userID,
				ProductID:   req.ProductID,
				ProductName: req.ProductName,
				Price:       req.Price,
				Quantity:    req.Quantity,
				ImageUrl:    req.ImageUrl,
			}
			if err := s.cartRepo.Create(newItem); err != nil {
				return nil, err
			}
			return newItem, nil
		}
		return nil, err
	}

	// Update quantity
	item.Quantity += req.Quantity
	if err := s.cartRepo.Update(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *CartService) UpdateQuantity(userID uint, req *models.UpdateCartQuantityRequest) (*models.CartItem, error) {
	item, err := s.cartRepo.FindByUserIDAndProductID(userID, req.ProductID)
	if err != nil {
		return nil, err
	}

	if req.Quantity <= 0 {
		if err := s.cartRepo.DeleteByProductID(userID, req.ProductID); err != nil {
			return nil, err
		}
		return nil, nil
	}

	item.Quantity = req.Quantity
	if err := s.cartRepo.Update(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *CartService) RemoveItem(userID, productID uint) error {
	return s.cartRepo.DeleteByProductID(userID, productID)
}

func (s *CartService) ClearCart(userID uint) error {
	return s.cartRepo.DeleteAllByUserID(userID)
}
