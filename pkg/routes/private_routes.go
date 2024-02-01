package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oyesaurav/go-todo/app/controllers"
	"github.com/oyesaurav/go-todo/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", middleware.JWTProtected(), controllers.Home) // create a new book

	// Routes for PUT method:
	route.Put("/book", middleware.JWTProtected(), controllers.Home) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book", middleware.JWTProtected(), controllers.Home) // delete one book by ID
}
