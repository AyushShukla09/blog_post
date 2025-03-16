package api

import (
	"blog_post/db"
	"blog_post/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// @Summary lists all blogs
// @Description Endpoint to list all blog posts
// @Tags Blogs
// @Produce json
// @Success 200 {array} models.Blog "Successful Response"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Router /blog-posts [get]
func GetAllBlogs(c *fiber.Ctx) error {
	blogs, err := db.DB.GetAllBlogs()
	if err != nil {
		log.Errorf("GetAllBlogs failed: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(blogs)
}

// @Summary fetch a blog
// @Description Endpoint to fetch a blog by id
// @Tags Blog
// @Produce json
// @Param id path int64 true "Blog ID"
// @Success 200 {object} models.Blog "Successful Response"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /blog-post/{id} [get]
func GetBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	blogID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("GetBlog failed: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	blog, err := db.DB.GetBlog(blogID)
	if err != nil {
		log.Errorf("GetBlog failed: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(blog)
}

// @Summary create a blog
// @Description Endpoint to create a blog by id
// @Tags Blog
// @Produce json
// @Accept json
// @Param request body models.BlogRequestBody true "Blog Request Body"
// @Success 201 {object} models.Blog "Successful Response"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /blog-post [post]
func CreateBlog(c *fiber.Ctx) error {
	var reqBody models.BlogRequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		log.Errorf("CreateBlog failed: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	blog, err := db.DB.CreateBlog(reqBody)
	if err != nil {
		log.Errorf("CreateBlog failed: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(blog)
}

// @Summary update a blog
// @Description Endpoint to update a blog by id
// @Tags Blog
// @Produce json
// @Accept json
// @Param id path int64 true "Blog ID"
// @Param request body models.BlogRequestBody true "Blog Request Body"
// @Success 200 {object} models.Blog "Successful Response"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /blog-post/{id} [put]
func UpdateBlog(c *fiber.Ctx) error {
	var reqBody models.BlogRequestBody
	id := c.Params("id")
	blogID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("UpdateBlog failed: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := c.BodyParser(&reqBody); err != nil {
		log.Errorf("UpdateBlog failed: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	blog, err := db.DB.UpdateBlog(blogID, reqBody)
	if err != nil {
		log.Errorf("UpdateBlog failed: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(blog)
}

// @Summary delete a blog
// @Description Endpoint to delete a blog by id
// @Tags Blog
// @Produce json
// @Param id path int64 true "Blog ID"
// @Success 200 {object} string "Successful Response"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /blog-post/{id} [delete]
func DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	blogID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("DeleteBlog failed: %v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := db.DB.DeleteBlog(blogID); err != nil {
		log.Errorf("DeleteBlog failed: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Blog deleted successfully"})
}
