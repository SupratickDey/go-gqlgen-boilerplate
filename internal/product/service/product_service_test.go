package service

import (
	"context"
	"testing"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
	"github.com/SupratickDey/go-gqlgen-boilerplate/internal/product/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Setup
var (
	productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productUseCase    = Product{ProductRepository: productRepository}
	name              = "Lg"
	stock             = 10
	products          = []*model.Product{
		{
			ID:    "1",
			Name:  &name,
			Stock: &stock,
		},
		{
			ID:    "2",
			Name:  &name,
			Stock: &stock,
		},
	}
)

func TestProductUseCase_GetProducts(t *testing.T) {
	ctx := context.Background()

	productRepository.Mock.On("GetProducts", ctx).Return(products)

	result := productUseCase.GetProducts(ctx)
	assert.NotNil(t, result)
}

func TestProductUseCase_GetSingleProduct(t *testing.T) {
	ctx := context.Background()

	productRepository.Mock.On("GetSingleProduct", ctx).Return(products[0])

	result := productUseCase.GetSingleProduct(ctx)
	assert.NotNil(t, result)
}
