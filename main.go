package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db "github.com/oBonn14/go-fiber-hex/config"
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

	_, err = db.NewDB(config.DB)
	if err != nil {
		os.Exit(1)
	}

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}]:${port} | ${status} | ${method} ${path} | ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello Obonn")
	})

	listenPort := fmt.Sprintf(":%s", config.HTTP.Port)
	app.Listen(listenPort)
}
