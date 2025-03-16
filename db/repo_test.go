package db

import (
	"blog_post/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRandomBlog(t *testing.T, r *Repo) models.Blog {
	blog, err := r.CreateBlog(models.BlogRequestBody{
		Title:       "Random Blog",
		Description: "Random Description",
		Body:        "Random Body",
	})
	assert.NoError(t, err)
	return blog
}

func TestGetAllBlogs(t *testing.T) {
	r := &Repo{
		data: make(map[int64]models.Blog),
	}
	t.Run("No Blogs in DB", func(t *testing.T) {
		blogs, err := r.GetAllBlogs()
		assert.Error(t, err)
		assert.Equal(t, errors.New("no blogs in DB"), err)
		assert.Empty(t, blogs)
	})
	t.Run("Successful Blog Creation", func(t *testing.T) {
		newBlog := createRandomBlog(t, r)
		blogs, err := r.GetAllBlogs()
		assert.NoError(t, err)
		assert.Equal(t, newBlog, blogs[0])
	})
}

func TestCreateBlog(t *testing.T) {
	r := &Repo{
		data: make(map[int64]models.Blog),
	}
	t.Run("Missing Required Fields", func(t *testing.T) {

		blog, err := r.CreateBlog(models.BlogRequestBody{
			Title:       "Random Blog",
			Description: "Random Description",
			Body:        "",
		})
		assert.Error(t, err, "missing required field")
		assert.Empty(t, blog)
	})
	t.Run("Successful Creation", func(t *testing.T) {
		blog := createRandomBlog(t, r)
		assert.NotEmpty(t, blog)
	})
}

func TestDeleteBlog(t *testing.T) {
	r := &Repo{
		data: make(map[int64]models.Blog),
	}
	t.Run("Blog Not Found", func(t *testing.T) {
		err := r.DeleteBlog(1000)
		assert.Error(t, err, "blog not found")
	})
	t.Run("Successful Deletion", func(t *testing.T) {
		blog := createRandomBlog(t, r)
		err := r.DeleteBlog(blog.ID)
		assert.NoError(t, err)
	})

}

func TestUpdateBlog(t *testing.T) {
	r := &Repo{
		data: make(map[int64]models.Blog),
	}
	t.Run("Missing Required Fields", func(t *testing.T) {

		blog, err := r.UpdateBlog(1000, models.BlogRequestBody{
			Title:       "Random Blog",
			Description: "Random Description",
			Body:        "",
		})
		assert.Error(t, err, "missing required field")
		assert.Empty(t, blog)
	})
	t.Run("Blog Not Found", func(t *testing.T) {
		blog, err := r.UpdateBlog(1000, models.BlogRequestBody{
			Title:       "Updated Blog",
			Description: "Updated Description",
			Body:        "Updated Body",
		})
		assert.Error(t, err, "blog not found")
		assert.Empty(t, blog)
	})
	t.Run("Successful Updation", func(t *testing.T) {
		blog := createRandomBlog(t, r)
		updatedBlog, err := r.UpdateBlog(blog.ID, models.BlogRequestBody{
			Title:       "Updated Blog",
			Description: "Updated Description",
			Body:        "Updated Body",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, blog)
		assert.NotEqual(t, blog.Body, updatedBlog.Body)
		assert.NotEqual(t, blog.Description, updatedBlog.Description)
		assert.NotEqual(t, blog.Title, updatedBlog.Title)
	})
}

func TestGetBlog(t *testing.T) {
	r := &Repo{
		data: make(map[int64]models.Blog),
	}
	t.Run("Blog Not Found", func(t *testing.T) {
		blog, err := r.GetBlog(1000)
		assert.Error(t, err)
		assert.Equal(t, errors.New("blog not found"), err)
		assert.Empty(t, blog)
	})
	t.Run("Successful Blog Creation", func(t *testing.T) {
		newBlog := createRandomBlog(t, r)
		blogs, err := r.GetBlog(newBlog.ID)
		assert.NoError(t, err)
		assert.Equal(t, newBlog, blogs)
	})
}
