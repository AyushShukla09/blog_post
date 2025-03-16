package middleware

import (
	"blog_post/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// VerifyBlogFields checks if required fields are present in the request body
func VerifyBlogFields(c *fiber.Ctx) error {
	var reqBody models.BlogRequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		log.Errorf("VerifyBlogFields failed: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}
	if reqBody.Body == "" || reqBody.Title == "" || reqBody.Description == "" {
		log.Error("Missing field")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required field",
		})
	}

	return c.Next()
}
