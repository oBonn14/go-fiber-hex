package port

import (
	"context"

	"github.com/oBonn14/go-fiber-hex/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product model.Product) (*model.Product, error)

	GetProduct(ctx context.Context, id string) (*model.Product, error)

	GetProducts(ctx context.Context) ([]*model.Product, error)

	UpdateProduct(ctx context.Context, id string, product model.Product) (*model.Product, error)

	DeleteProduct(ctx context.Context, id string) (*model.Product, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, product model.Product) (*model.Product, error)

	GetProduct(ctx context.Context, id string) (*model.Product, error)

	GetProducts(ctx context.Context) ([]*model.Product, error)

	UpdateProduct(ctx context.Context, id string, product model.Product) (*model.Product, error)

	DeleteProduct(ctx context.Context, id string) (*model.Product, error)
}
