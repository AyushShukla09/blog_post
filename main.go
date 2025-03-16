package main

import (
	"blog_post/api"

	_ "blog_post/docs"
	m "blog_post/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/spf13/viper"
)

//	@title	Blog API

// @version		1.0
// @contact.name	Ayush Shukla
// @contact.email	ayush.shukla8797@gmail.com
// @host			quartiz-blog-post.onrender.com
// @BasePath		/api/v1
// @Schemes https
func main() {
	initConfig()
	siteURL := viper.GetString("SITE_URL")
	if siteURL == "" {
		log.Error("SITE_URL is not set. Please configure it in .env file or environment variables.")
	}
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}

	app := setup()
	log.Info("Listening at port: " + siteURL + port)
	log.Fatal(app.Listen(":" + port))
}

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file: %v", err)
	}
}

func setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders: "*",
	}))
	router := app.Group("/api/v1")
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	router.Get("/blog-posts", api.GetAllBlogs)
	router.Post("/blog-post", m.VerifyBlogFields, api.CreateBlog)
	router.Get("/blog-post/:id<min(1)>", api.GetBlog)
	router.Put("/blog-post/:id<min(1)>", api.UpdateBlog)
	router.Delete("/blog-post/:id<min(1)>", api.DeleteBlog)

	return app
}
