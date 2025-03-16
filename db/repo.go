package db

import (
	"blog_post/models"
	"errors"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Repo struct {
	data map[int64]models.Blog
	mu   sync.RWMutex
}

var DB = Repo{
	data: make(map[int64]models.Blog),
}

// GetAllBlogs lists all blogs
func (r *Repo) GetAllBlogs() ([]models.Blog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	blogs := make([]models.Blog, 0, len(r.data))
	for _, blog := range r.data {
		blogs = append(blogs, blog)
	}
	if len(blogs) == 0 {
		return nil, errors.New("no blogs in DB")
	}
	return blogs, nil
}

// GetBlog fetches a blog by id
func (r *Repo) GetBlog(id int64) (models.Blog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	blog, exists := r.data[id]
	if !exists {
		return models.Blog{}, errors.New("blog not found")
	}
	return blog, nil
}

// DeleteBlog deletes a blog
func (r *Repo) DeleteBlog(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.data[id]
	if !exists {
		return errors.New("blog not found")
	}
	delete(r.data, id)
	return nil
}

// CreateBlog creates a new blog
func (r *Repo) CreateBlog(blog models.BlogRequestBody) (models.Blog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if blog.Title == "" || blog.Body == "" || blog.Description == "" {
		log.Error("CreateBlog failed: Missing field")
		return models.Blog{}, errors.New("missing required field")
	}
	id := len(r.data)
	newID := int64(id + 1)
	r.data[newID] = models.Blog{
		ID:          newID,
		Title:       blog.Title,
		Description: blog.Description,
		Body:        blog.Body,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return r.data[newID], nil
}

// UpdateBlog updates an existing blog
func (r *Repo) UpdateBlog(id int64, blog models.BlogRequestBody) (models.Blog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if blog.Title == "" || blog.Body == "" || blog.Description == "" {
		log.Error("UpdateBlog failed: Missing field")
		return models.Blog{}, errors.New("missing required field")
	}
	oldBlog, exists := r.data[id]
	if !exists {
		return models.Blog{}, errors.New("blog not found")
	}
	newBlog := models.Blog{
		ID:          id,
		Title:       blog.Title,
		Description: blog.Description,
		Body:        blog.Body,
		UpdatedAt:   time.Now(),
		CreatedAt:   oldBlog.CreatedAt,
	}
	r.data[id] = newBlog
	return newBlog, nil
}
