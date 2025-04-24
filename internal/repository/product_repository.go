package repository

import (
	"github.com/rizdian/go-full-api/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	FindAll() ([]model.Product, error)
	BeginTransaction() *gorm.DB
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
