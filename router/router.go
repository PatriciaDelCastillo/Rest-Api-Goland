package router

import (
	"github.com/PatriciaDelCastillo/Rest-Api-Goland/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/hello", handlers.Hello)
	app.Get("/get", handlers.GetTasks)
	app.Get("/get/:id",handlers.CreateTasksId)
	app.Post("/post", handlers.CreateTasks)
	app.Put("/editar/:id", handlers.UpdateTasks)
	app.Delete("/delete/:id", handlers.DeleteTask)
}