package service

import (
	"errors"
	"github.com/rizdian/go-full-api/internal/model"
	"github.com/rizdian/go-full-api/internal/repository"
	"gorm.io/gorm/clause"
)

type OrderService interface {
	Create(order *model.Order) error
	GetAll() ([]model.Order, error)
}

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{orderRepo: orderRepo, productRepo: productRepo}
}

func (s *orderService) Create(order *model.Order) error {
	tx := s.productRepo.BeginTransaction()

	// Fetch product with pessimistic lock
	var product model.Product
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, order.ProductID).Error; err != nil {
		tx.Rollback()
		return errors.New("product not found")
	}

	// Check if enough stocks are available
	if product.Stock < order.Quantity {
		tx.Rollback()
		return errors.New("not enough stock available")
	}

	// Update product stock
	product.Stock -= order.Quantity
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Calculate total price
	order.Total = float64(order.Quantity) * product.Price

	// Save order to DB
	if err := s.orderRepo.Create(order); err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

func (s *orderService) GetAll() ([]model.Order, error) {
	// Fetch all orders from DB
	orders, err := s.orderRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to fetch orders")
	}
	return orders, nil
}
