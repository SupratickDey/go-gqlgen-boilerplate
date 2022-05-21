package repository

import (
	"context"
	"github.com/SupratickDey/go-gqlgen-boilerplate/pkg/database"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

type ProductRepositoryInterface interface {
	GetProducts(ctx context.Context) []*model.Product
	GetSingleProduct(ctx context.Context) *model.Product
}

type ProductRepository struct {
	db *database.Conn
}

func NewProductRepository(db *database.Conn) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (pr ProductRepository) GetProducts(ctx context.Context) []*model.Product {
	var products []*model.Product

	//pr.db.Find(&products)

	return products
}

func (pr ProductRepository) GetSingleProduct(ctx context.Context) *model.Product {
	product := &model.Product{}

	//pr.db.First(product)

	return product
}
