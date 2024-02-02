package main

import (
	_ "fmt"

	"github.com/oyesaurav/go-todo/pkg/configs"
	"github.com/oyesaurav/go-todo/pkg/middleware"
	"github.com/oyesaurav/go-todo/pkg/routes"
	"github.com/oyesaurav/go-todo/pkg/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	utils.CronScheduler()
	utils.StartServer(app)
}
