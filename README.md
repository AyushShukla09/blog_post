# blog_post

This is a RESTful API for managing blog posts, built with Go.

## Features

* **CRUD operations for blog posts:** 
  * Create new posts
  * Retrieve all posts
  * Get a post by ID
  * Update a post
  * Delete a post 
* **Database:** In Memory
* **Framework:** Fibre(v2)

## Project Structure

```
/blog_post
│
├── /api             # Handlers and routes
├── /db              # In-memory database implementation
├── /models          # Models for request/response structures
├── main.go          # Application entry point
├── go.mod           # Go module file
└── go.sum           # Dependencies file

```

## Routes

```
POST    /api/blog-post — Add a blog post
GET     /api/blog-post — Get all blog posts
GET     /api/blog-post/:id — Get single blog post
DELETE  /api/blog-post/:id — Delete a blog post
PATCH   /api/blog-post/:id — Update a blog post
```