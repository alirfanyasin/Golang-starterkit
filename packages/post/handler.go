package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetPosts godoc
// @Summary      List all blog posts
// @Description  Get a list of all blog posts
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string][]Post
// @Router       /posts [get]
func (h *Handler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GetPost godoc
// @Summary      Get a single blog post
// @Description  Get detail of a blog post by ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  map[string]Post
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /posts/{id} [get]
func (h *Handler) GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// CreatePost godoc
// @Summary      Create a new blog post
// @Description  Create a blog post with the given details
// @Tags         posts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        post body      CreatePostInput  true  "Create Post Input"
// @Success      201  {object}  map[string]Post
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      403  {object}  map[string]string
// @Router       /posts [post]
func (h *Handler) CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

// UpdatePost godoc
// @Summary      Update an existing blog post
// @Description  Update details of a blog post by ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Post ID"
// @Param        post body      UpdatePostInput  true  "Update Post Input"
// @Success      200  {object}  map[string]Post
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      403  {object}  map[string]string
// @Router       /posts/{id} [put]
func (h *Handler) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.service.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DeletePost godoc
// @Summary      Delete a blog post
// @Description  Remove a blog post by ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      403  {object}  map[string]string
// @Router       /posts/{id} [delete]
func (h *Handler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
