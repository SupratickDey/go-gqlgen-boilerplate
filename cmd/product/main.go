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

	brand_service "github.com/SupratickDey/go-gqlgen-boilerplate/controllers/brand/service"
	prod_service "github.com/SupratickDey/go-gqlgen-boilerplate/controllers/product/service"

	brand_repo "github.com/SupratickDey/go-gqlgen-boilerplate/controllers/brand/repository"
	prod_repo "github.com/SupratickDey/go-gqlgen-boilerplate/controllers/product/repository"
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

func productResolverInit(db *database.Conn) prod_service.Product {
	productRepo := prod_repo.NewProductRepository(db)
	return prod_service.Product{
		ProductRepository: productRepo,
	}
}

func brandResolverInit(db *database.Conn) brand_service.Brand {
	brandRepo := brand_repo.NewBrandRepository(db)
	return brand_service.Brand{
		BrandRepository: brandRepo,
	}
}
