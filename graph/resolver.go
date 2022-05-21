package graph

import (
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/dataloader"
	"github.com/SupratickDey/go-gqlgen-boilerplate/internal/brand"
	"github.com/SupratickDey/go-gqlgen-boilerplate/internal/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductResolver product.ProductInterface
	BrandResolver   brand.BrandInterface
	DataLoaders     dataloader.Retriever
}
