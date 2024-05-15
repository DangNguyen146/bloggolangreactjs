package routes

import (
	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/controller"
	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthorized)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/allpost/:id", controller.UpdatePost)
	app.Post("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload-image", controller.Upload)
	app.Static("/api/images", "./uploads")
}
