package service

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/controllers/brand/repository"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

// brand service
type Brand struct {
	BrandRepository repository.BrandRepositoryInterface
}

func NewProductResolver(productRepo repository.BrandRepositoryInterface) Brand {
	return Brand{
		BrandRepository: productRepo,
	}
}

func (b Brand) GetBrands(ctx context.Context) []*model.Brand {
	return b.BrandRepository.GetBrands(ctx)
}
