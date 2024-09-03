package service

import (
	"context"
	"github.com/oBonn14/go-fiber-hex/model"
	"github.com/oBonn14/go-fiber-hex/port"
)

type ProductService struct {
	repo port.ProductRepositoryInterface
}

func NewProductService(repo port.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		repo,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	product, err := ps.repo.CreateProduct(ctx, *product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	product, err := ps.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := ps.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) UpdateProduct(ctx context.Context, id string, product model.Product) (*model.Product, error) {
	produuctUpdate, err := ps.repo.UpdateProduct(ctx, id, product)
	if err != nil {
		return nil, err
	}
	return produuctUpdate, nil
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id string) (*model.Product, error) {
	product, err := ps.repo.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
