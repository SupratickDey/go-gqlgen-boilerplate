package brand

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

type BrandInterface interface {
	GetBrands(ctx context.Context) []*model.Brand
}
