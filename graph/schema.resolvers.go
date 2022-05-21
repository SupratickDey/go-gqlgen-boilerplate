package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/generated"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/model"
)

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) CreateProduct(ctx context.Context, data model.CreateProductInput) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProducts(ctx context.Context) ([]*model.Product, error) {
	return r.ProductResolver.GetProducts(ctx), nil
}

func (r *queryResolver) GetProduct(ctx context.Context) (*model.Product, error) {
	return r.ProductResolver.GetSingleProduct(ctx), nil
}

func (r *queryResolver) GetBrands(ctx context.Context) ([]*model.Brand, error) {
	return r.BrandResolver.GetBrands(ctx), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
