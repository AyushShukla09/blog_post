package api

import (
	"blog_post/models"
	"fmt"
	"io"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func CreateRandomBlog(t *testing.T) models.Blog {
	app := fiber.New()
	app.Post("/blog-post", CreateBlog)
	var createdBlog models.Blog
	t.Run("Successful creation", func(t *testing.T) {
		reqBody := models.BlogRequestBody{
			Title:       "Test Title",
			Description: "Test Description",
			Body:        "Test Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/blog-post", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &createdBlog)
	})
	return createdBlog
}
func TestGetAllBlogs(t *testing.T) {
	app := fiber.New()
	app.Get("/blog-posts", GetAllBlogs)
	t.Run("No blogs in DB", func(t *testing.T) {
		expectedBlogs := []models.Blog{}
		req := httptest.NewRequest(http.MethodGet, "/blog-posts", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, len(expectedBlogs), len(actualBlogs))
	})
	t.Run("Successful retrieval", func(t *testing.T) {
		expectedBlogs := CreateRandomBlog(t)
		req := httptest.NewRequest("GET", "/blog-posts", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, expectedBlogs.ID, actualBlogs[0].ID)
		assert.Equal(t, expectedBlogs.Title, actualBlogs[0].Title)
		assert.Equal(t, expectedBlogs.Description, actualBlogs[0].Description)
		assert.Equal(t, expectedBlogs.Body, actualBlogs[0].Body)
	})
}

func TestCreateBlog(t *testing.T) {
	app := fiber.New()
	app.Post("/blog-post", CreateBlog)
	var createdBlog models.Blog
	t.Run("Successful creation", func(t *testing.T) {
		reqBody := models.BlogRequestBody{
			Title:       "Test Title",
			Description: "Test Description",
			Body:        "Test Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/blog-post", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &createdBlog)
		assert.Equal(t, reqBody.Title, createdBlog.Title)
		assert.Equal(t, reqBody.Description, createdBlog.Description)
		assert.Equal(t, reqBody.Body, createdBlog.Body)
	})
	t.Run("Invalid JSON format", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/blog-post", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Empty request body", func(t *testing.T) {
		reqBody := models.BlogRequestBody{
			Title:       "",
			Description: "",
			Body:        "Test Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/blog-post", bytes.NewReader([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		response, _ := io.ReadAll(resp.Body)
		var result map[string]string
		json.Unmarshal(response, &result)
		assert.Contains(t, result["error"], "missing required field")
	})
}

func TestGetBlog(t *testing.T) {
	app := fiber.New()
	app.Get("/blog-post/:id", GetBlog)
	t.Run("No blogs in DB", func(t *testing.T) {
		expectedBlogs := []models.Blog{}
		req := httptest.NewRequest("GET", fmt.Sprintf("/blog-post/%d", 1000), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, len(expectedBlogs), len(actualBlogs))
	})
	t.Run("Invalid Blog ID", func(t *testing.T) {
		expectedBlogs := []models.Blog{}
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/blog-post/%s", "a"), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, len(expectedBlogs), len(actualBlogs))
	})
	t.Run("Successful retrieval", func(t *testing.T) {
		expectedBlogs := CreateRandomBlog(t)
		u, _ := json.Marshal(expectedBlogs)
		fmt.Println(string(u))
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/blog-post/%d", expectedBlogs.ID), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var actualBlogs models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, expectedBlogs.ID, actualBlogs.ID)
		assert.Equal(t, expectedBlogs.Title, actualBlogs.Title)
		assert.Equal(t, expectedBlogs.Description, actualBlogs.Description)
		assert.Equal(t, expectedBlogs.Body, actualBlogs.Body)
	})
}

func TestDeleteBlog(t *testing.T) {
	app := fiber.New()
	app.Delete("/blog-post/:id", DeleteBlog)
	t.Run("Invalid Blog ID", func(t *testing.T) {
		expectedBlogs := []models.Blog{}
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/blog-post/%s", "a"), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, len(expectedBlogs), len(actualBlogs))
	})
	t.Run("No blogs in DB", func(t *testing.T) {
		expectedBlogs := []models.Blog{}
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/blog-post/%d", 1000), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, len(expectedBlogs), len(actualBlogs))
	})
	t.Run("Successful Deleteion", func(t *testing.T) {
		expectedBlogs := CreateRandomBlog(t)
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/blog-post/%d", expectedBlogs.ID), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		response, _ := io.ReadAll(resp.Body)
		fmt.Println(string(response))
		assert.Contains(t, string(response), "Blog deleted successfully")
	})
}

func TestUpdateBlog(t *testing.T) {
	app := fiber.New()
	app.Put("/blog-post/:id", UpdateBlog)
	t.Run("Invalid Blog ID", func(t *testing.T) {
		expectedBlogs := []models.Blog{}
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/blog-post/%s", "a"), nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var actualBlogs []models.Blog
		json.NewDecoder(resp.Body).Decode(&actualBlogs)
		assert.Equal(t, len(expectedBlogs), len(actualBlogs))
	})
	var createdBlog models.Blog
	t.Run("Successful Updation", func(t *testing.T) {
		expectedBlogs := CreateRandomBlog(t)
		reqBody := models.BlogRequestBody{
			Title:       "Test Updated Title",
			Description: "Test Updated Description",
			Body:        "Test Updated Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/blog-post/%d", expectedBlogs.ID), bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &createdBlog)
		assert.Equal(t, reqBody.Title, createdBlog.Title)
		assert.Equal(t, reqBody.Description, createdBlog.Description)
		assert.Equal(t, reqBody.Body, createdBlog.Body)
	})
	t.Run("Invalid JSON format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/blog-post/%d", 1000), bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Empty request body", func(t *testing.T) {
		reqBody := models.BlogRequestBody{
			Title:       "",
			Description: "",
			Body:        "Test Body",
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/blog-post/%d", 1000), bytes.NewReader([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		response, _ := io.ReadAll(resp.Body)
		var result map[string]string
		json.Unmarshal(response, &result)
		assert.Contains(t, result["error"], "missing required field")
	})
}
