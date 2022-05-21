package repository

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/pkg/database"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

type BrandRepositoryInterface interface {
	GetBrands(ctx context.Context) []*model.Brand
}

type BrandRepository struct {
	db *database.Conn
}

func NewBrandRepository(db *database.Conn) BrandRepositoryInterface {
	return &BrandRepository{
		db: db,
	}
}

func (br BrandRepository) GetBrands(ctx context.Context) []*model.Brand {
	var brands []*model.Brand

	// use sql query to fetch brands br.db.Preload("Products").Find(&brands)

	return brands
}
