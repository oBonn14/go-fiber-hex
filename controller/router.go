package controller

import "github.com/gofiber/fiber/v2"

type Router struct {
	*fiber.App
}

func NewRouter(app *fiber.App, controller ProductController) {
	app.Post("/addProduct", controller.CreateProduct)
	app.Get("/product/:id", controller.GetProduct)
	app.Get("/product", controller.GetProducts)
	app.Put("/product/:id", controller.UpdateProduct)
	app.Delete("/product/:id", controller.DeleteProduct)
}
