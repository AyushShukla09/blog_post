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
├── /docs            # Swagger documentation
├── /middlewares     # Middlewares for Request
├── /models          # Models for request/response structures
├── main_test.go     # Testing main file
├── main.go          # Application entry point
├── Makefile         # Makefile to run commands
├── go.mod           # Go module file
└── go.sum           # Dependencies file

```

## Routes

```
POST    /api/blog-post     — Add a blog post
GET     /api/blog-posts    — Get all blog posts
GET     /api/blog-post/:id — Get single blog post
DELETE  /api/blog-post/:id — Delete a blog post
PATCH   /api/blog-post/:id — Update a blog post
```

## Swagger Link

https://quartiz-blog-post.onrender.com/swagger/index.html

## Commands

At the root of project, i.e., where main.go is present, run below mentioned commands**
*  **make swag:** To regenrate swagger docs
*  **make test:** To unit test whole application
*  **make run:** To run the application 

## Test Coverage

```
go test --cover ./...
ok      blog_post       0.003s  coverage: 75.0% of statements
ok      blog_post/api   0.004s  coverage: 100.0% of statements
ok      blog_post/db    0.002s  coverage: 100.0% of statements
        blog_post/docs          coverage: 0.0% of statements
ok      blog_post/middlewares   (cached)        coverage: 100.0% of statements
?       blog_post/models        [no test files]
```
