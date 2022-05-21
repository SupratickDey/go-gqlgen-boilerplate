package repository

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repo *ProductRepositoryMock) GetProducts(ctx context.Context) []*model.Product {
	products := []*model.Product{}
	return products
}

func (repo *ProductRepositoryMock) GetSingleProduct(ctx context.Context) *model.Product {
	arguments := repo.Mock.Called(ctx)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(*model.Product)
	return product
}
