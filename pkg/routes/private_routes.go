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
	route.Get("/task", middleware.JWTProtected(), controllers.GetTasks) 
	route.Get("/subtask", middleware.JWTProtected(), controllers.GetSubTasks)
	route.Post("/task", middleware.JWTProtected(), controllers.CreateTasks) 
	route.Post("/subtask", middleware.JWTProtected(), controllers.CreateSubTasks) 
	route.Patch("/task", middleware.JWTProtected(), controllers.UpdateTasks) 
	route.Patch("/subtask", middleware.JWTProtected(), controllers.UpdateSubTasks) 
	route.Delete("/task", middleware.JWTProtected(), controllers.DeleteTasks) 
	route.Delete("/subtask", middleware.JWTProtected(), controllers.DeleteSubTasks) 
}
