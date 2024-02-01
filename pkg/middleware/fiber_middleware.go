package middleware

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(cors.Config{
		    AllowOriginsFunc: func(origin string) bool {
				return os.Getenv("ENVIRONMENT") == "development"
		    },
	    }),
	logger.New(),
   )
}