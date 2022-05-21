package product

import (
	"context"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

//nolint:golint // ProductInterface is to indicate interface for products
type ProductInterface interface {
	GetProducts(ctx context.Context) []*model.Product
	GetSingleProduct(ctx context.Context) *model.Product
}
