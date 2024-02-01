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

	route.Get("", middleware.JWTProtected(), controllers.Home)

	// Routes for POST method:
	route.Post("/task", middleware.JWTProtected(), controllers.CreateTasks) // create a new book
	route.Post("/subtask", middleware.JWTProtected(), controllers.CreateSubTasks) // create a new book
	route.Patch("/task", middleware.JWTProtected(), controllers.UpdateTasks) 
	route.Patch("/subtask", middleware.JWTProtected(), controllers.UpdateSubTasks) 
	
}
