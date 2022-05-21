package dataloader

import (
	"context"
	"time"

	"github.com/SupratickDey/go-gqlgen-boilerplate/controllers/product/repository"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

type contextKey string

const key = contextKey("dataloaders")

type Loaders struct {
	ProductsByBrandID *ProductLoader
}

func NewLoader(ctx context.Context, repo repository.ProductRepositoryInterface) *Loaders {
	return &Loaders{
		ProductsByBrandID: newProductsByBrandID(ctx, repo),
	}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}

func newProductsByBrandID(ctx context.Context, repo repository.ProductRepositoryInterface) *ProductLoader {
	return NewProductLoader(ProductLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(brandIDs []string) ([]model.Product, []error) {
			res := repo.GetProducts(ctx)

			groupByBrandID := make(map[string]model.Product, len(brandIDs))
			for _, r := range res {
				groupByBrandID[*r.BrandID] = model.Product{
					ID:    r.ID,
					Name:  r.Name,
					Stock: r.Stock,
				}
			}

			result := make([]model.Product, len(brandIDs))
			for i, brandID := range brandIDs {
				result[i] = groupByBrandID[brandID]
			}

			return result, nil
		},
	})
}
