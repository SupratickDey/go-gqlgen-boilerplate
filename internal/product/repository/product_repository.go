package repository

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryInterface interface {
	GetProducts(ctx context.Context) []*model.Product
	GetSingleProduct(ctx context.Context) *model.Product
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (pr ProductRepository) GetProducts(ctx context.Context) []*model.Product {
	var products []*model.Product

	pr.db.Find(&products)

	return products
}

func (pr ProductRepository) GetSingleProduct(ctx context.Context) *model.Product {
	product := &model.Product{}

	pr.db.First(product)

	return product
}
