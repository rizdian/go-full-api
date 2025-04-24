package service

import (
	"errors"
	"github.com/rizdian/go-full-api/internal/model"
	"github.com/rizdian/go-full-api/internal/repository"
)

type ProductService interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) Create(product *model.Product) error {
	// Validate stock
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Save product to DB
	return s.productRepo.Create(product)
}

func (s *productService) GetAll() ([]model.Product, error) {
	// Fetch all products from DB
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to fetch products")
	}
	return products, nil
}
