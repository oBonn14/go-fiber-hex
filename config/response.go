package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oBonn14/go-fiber-hex/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

type productResponse struct {
	ID      primitive.ObjectID `json:"id"`
	Product string             `json:"product"`
	Stock   int                `json:"stock"`
}

func NewProductResponse(product *model.Product) productResponse {
	return productResponse{
		ID:      product.ID,
		Product: product.Product,
		Stock:   product.Stock,
	}
}

func HandleResponse(ctx *fiber.Ctx, message string, data any) {
	rsp := newResponse(true, message, data)
	ctx.JSON(rsp)
}
