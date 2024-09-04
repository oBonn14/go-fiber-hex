package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db "github.com/oBonn14/go-fiber-hex/config"
	"github.com/oBonn14/go-fiber-hex/controller"
	"github.com/oBonn14/go-fiber-hex/repository"
	"github.com/oBonn14/go-fiber-hex/service"
	"log/slog"
	"os"
)

func main() {
	config, err := db.New()

	if err != nil {
		slog.Error("Error opening database: ", err)
		os.Exit(1)
	}

	db.Set(config.App)

	slog.Info("Database app initialized", "app", config.App.Name, "env", config.App.Env)

	data, _ := db.NewDB(config.DB)
	repo := repository.NewProductRepository(data)
	serv := service.NewProductService(repo)
	productController := controller.NewProductController(serv)

	//if err != nil {
	//	os.Exit(1)
	//}

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}]:${port} | ${status} | ${method} ${path} | ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	controller.NewRouter(app, *productController)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello Obonn")
	})

	listenPort := fmt.Sprintf(":%s", config.HTTP.Port)
	app.Listen(listenPort)
}
