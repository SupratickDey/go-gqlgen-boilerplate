package product

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

type ProductInterface interface {
	GetProducts(ctx context.Context) []*model.Product
	GetSingleProduct(ctx context.Context) *model.Product
}
