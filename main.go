package main

import (
	"blog_post/api"

	_ "blog_post/docs"
	m "blog_post/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/swagger" // swagger handler
)

//	@title	Blog API

// @version		1.0
// @contact.name	Ayush Shukla
// @contact.email	ayush.shukla8797@gmail.com
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	app := setup()
	log.Info("Listening at port :8080")
	log.Fatal(app.Listen(":8080"))
}

func setup() *fiber.App {
	app := fiber.New()
	router := app.Group("/api/v1")
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	router.Get("/blog-posts", api.GetAllBlogs)
	router.Post("/blog-post", m.VerifyBlogFields, api.CreateBlog)
	router.Get("/blog-post/:id<min(1)>", api.GetBlog)
	router.Put("/blog-post/:id<min(1)>", api.UpdateBlog)
	router.Delete("/blog-post/:id<min(1)>", api.DeleteBlog)

	return app
}
