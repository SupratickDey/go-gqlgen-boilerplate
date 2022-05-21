package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/dataloader"
	"github.com/SupratickDey/go-gqlgen-boilerplate/graph/generated"
	"github.com/SupratickDey/go-gqlgen-boilerplate/pkg/database"

	brand_usecase "github.com/SupratickDey/go-gqlgen-boilerplate/internal/brand/service"
	prod_usecase "github.com/SupratickDey/go-gqlgen-boilerplate/internal/product/service"

	brand_repo "github.com/SupratickDey/go-gqlgen-boilerplate/internal/brand/repository"
	prod_repo "github.com/SupratickDey/go-gqlgen-boilerplate/internal/product/repository"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbConn := database.DBInit()
	dl := dataloader.NewRetriever()

	var resolver = &graph.Resolver{
		ProductResolver: productResolverInit(dbConn),
		BrandResolver:   brandResolverInit(dbConn),
		DataLoaders:     dl,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func productResolverInit(db *database.Conn) prod_usecase.Product {
	productRepo := prod_repo.NewProductRepository(db)
	return prod_usecase.Product{
		ProductRepository: productRepo,
	}
}

func brandResolverInit(db *database.Conn) brand_usecase.Brand {
	brandRepo := brand_repo.NewBrandRepository(db)
	return brand_usecase.Brand{
		BrandRepository: brandRepo,
	}
}
