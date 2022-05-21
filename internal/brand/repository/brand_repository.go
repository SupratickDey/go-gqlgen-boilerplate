package repository

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
	"github.com/jinzhu/gorm"
)

type BrandRepositoryInterface interface {
	GetBrands(ctx context.Context) []*model.Brand
}

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepositoryInterface {
	return &BrandRepository{
		db: db,
	}
}

func (br BrandRepository) GetBrands(ctx context.Context) []*model.Brand {
	var brands []*model.Brand

	br.db.Preload("Products").Find(&brands)

	return brands
}
