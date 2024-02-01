package routes

import(
	"github.com/gofiber/fiber/v2"
	"github.com/oyesaurav/go-todo/app/controllers"
)

func PublicRoutes(app *fiber.App){
	auth := app.Group("/auth")
	auth.Get("/login", controllers.Login)
	auth.Get("/callback", controllers.Callback)


	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", controllers.Home)

	// auth := v1.Group("/auth")
	// auth.Get("/", controllers.Home)
}