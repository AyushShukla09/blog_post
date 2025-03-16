package middleware

import (
	"blog_post/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestVerifyBlogFields(t *testing.T) {
	app := fiber.New()
	t.Run("Valid request body", func(t *testing.T) {
		reqBody := models.BlogRequestBody{
			Title:       "Test Title",
			Description: "Test Description",
			Body:        "Test Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/", VerifyBlogFields, func(c *fiber.Ctx) error {
			return c.SendString("OK")
		})

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
	t.Run("Invalid JSON format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/", VerifyBlogFields, func(c *fiber.Ctx) error {
			return c.SendString("OK")
		})

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		respBody, _ := io.ReadAll(resp.Body)
		var res map[string]any
		json.Unmarshal(respBody, &res)
		assert.Equal(t, "Invalid JSON format", res["error"])
	})
	t.Run("Missing title", func(t *testing.T) {
		reqBody := models.BlogRequestBody{
			Description: "Test Description",
			Body:        "Test Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/", VerifyBlogFields, func(c *fiber.Ctx) error {
			return c.SendString("OK")
		})

		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		respBody, _ := io.ReadAll(resp.Body)
		var res map[string]any
		json.Unmarshal(respBody, &res)
		assert.Equal(t, "Missing required field", res["error"])

	})
}
