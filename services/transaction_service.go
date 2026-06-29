package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/SadamIRM/be-smoke-store/models"
	"github.com/SadamIRM/be-smoke-store/repositories"
)

type TransactionService struct {
	txRepo   *repositories.TransactionRepository
	cartRepo *repositories.CartRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		txRepo:   repositories.NewTransactionRepository(),
		cartRepo: repositories.NewCartRepository(),
	}
}

func (s *TransactionService) Checkout(userID uint, paymentMethod string) (*models.Transaction, error) {
	// 1. Get cart items
	cartItems, err := s.cartRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(cartItems) == 0 {
		return nil, errors.New("keranjang belanja kosong")
	}

	// 2. Calculate total and map items
	var totalAmount float64
	var txItems []models.TransactionItem

	for _, cartItem := range cartItems {
		totalAmount += cartItem.Price * float64(cartItem.Quantity)
		txItems = append(txItems, models.TransactionItem{
			ProductID:   cartItem.ProductID,
			ProductName: cartItem.ProductName,
			Price:       cartItem.Price,
			Quantity:    cartItem.Quantity,
			ImageUrl:    cartItem.ImageUrl,
		})
	}

	// 3. Generate transaction number
	txNum := fmt.Sprintf("TRX-%d", time.Now().UnixMilli())

	status := "Selesai"
	if paymentMethod == "Smoke Money" {
		status = "Menunggu Pembayaran"
	}

	// 4. Create transaction struct
	tx := &models.Transaction{
		UserID:            userID,
		TransactionNumber: txNum,
		TotalAmount:       totalAmount,
		PaymentMethod:     paymentMethod,
		Status:            status,
		Items:             txItems,
	}

	// 5. Save to database
	if err := s.txRepo.Create(tx); err != nil {
		return nil, err
	}

	// 6. Clear cart
	if err := s.cartRepo.DeleteAllByUserID(userID); err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *TransactionService) GetTransactions(userID uint) ([]models.Transaction, error) {
	return s.txRepo.FindAllByUserID(userID)
}
