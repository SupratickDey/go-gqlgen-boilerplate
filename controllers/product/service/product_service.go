package service

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/controllers/product/repository"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

// product service
type Product struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewProductResolver(productRepo repository.ProductRepositoryInterface) Product {
	return Product{
		ProductRepository: productRepo,
	}
}

func (p Product) GetProducts(ctx context.Context) []*model.Product {
	return p.ProductRepository.GetProducts(ctx)
}

func (p Product) GetSingleProduct(ctx context.Context) *model.Product {
	return p.ProductRepository.GetSingleProduct(ctx)
}
