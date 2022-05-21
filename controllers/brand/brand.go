package brand

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

//nolint:golint // BrandInterface is to indicate interface for brands
type BrandInterface interface {
	GetBrands(ctx context.Context) []*model.Brand
}
