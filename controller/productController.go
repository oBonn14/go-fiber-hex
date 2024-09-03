package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oBonn14/go-fiber-hex/config"
	"github.com/oBonn14/go-fiber-hex/model"
	"github.com/oBonn14/go-fiber-hex/port"
)

type ProductController struct {
	ctr port.ProductServiceInterface
}

func NewProductController(ctr port.ProductServiceInterface) *ProductController {
	return &ProductController{
		ctr,
	}
}

type createProductRequest struct {
	ProductName string `json:"productName" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}

func (pc *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var productRequest createProductRequest
	if err := ctx.BodyParser(&productRequest); err != nil {
		config.HandleResponse(ctx, "Error", err)
		return err
	}

	product := model.Product{
		Product: productRequest.ProductName,
		Stock:   productRequest.Stock,
	}

	result, err := pc.ctr.CreateProduct(ctx.Context(), &product)
	if err != nil {
		config.HandleResponse(ctx, "Error", err)
		return err
	}

	rsp := config.NewProductResponse(result)
	config.HandleResponse(ctx, "Success", rsp)
	return nil
}

func (pc *ProductController) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		config.HandleResponse(ctx, "Error", "Parameter id null")
	}

	result, err := pc.ctr.GetProduct(ctx.Context(), id)
	if err != nil {
		return nil
	}

	rsp := config.NewProductResponse(result)
	config.HandleResponse(ctx, "Success Geting Product", rsp)
	return nil
}

func (pc *ProductController) GetProducts(ctx *fiber.Ctx) error {
	result, err := pc.ctr.GetProducts(ctx.Context())
	if err != nil {
		return err
	}
	config.HandleResponse(ctx, "Success Geting Products", result)
	return nil
}

type updateProductRequest struct {
	Product string `json:"product" validate:"required"`
	Stock   int    `json:"stock" validate:"required"`
}

func (pc *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		config.HandleResponse(ctx, "Error", "Parameter id null")
	}

	var productRequest updateProductRequest
	if err := ctx.BodyParser(&productRequest); err != nil {
		config.HandleResponse(ctx, "Error", err)
		return err
	}
	product := model.Product{
		Product: productRequest.Product,
		Stock:   productRequest.Stock,
	}

	result, err := pc.ctr.UpdateProduct(ctx.Context(), id, product)
	if err != nil {
		config.HandleResponse(ctx, "Error", err)
		return err
	}

	rsp := config.NewProductResponse(result)
	config.HandleResponse(ctx, "Success Updating Product", rsp)
	return nil
}

func (pc *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		config.HandleResponse(ctx, "Error", "Parameter id null")
	}

	result, err := pc.ctr.DeleteProduct(ctx.Context(), id)
	if err != nil {
		return err
	}
	rsp := config.NewProductResponse(result)
	config.HandleResponse(ctx, "Success Deleting Product", rsp)
	return nil

}
