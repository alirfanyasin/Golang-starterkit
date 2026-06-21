package routes

import (
	"github.com/alirfanyasin/golang-starterkit/config"
	"github.com/alirfanyasin/golang-starterkit/middleware"
	"github.com/alirfanyasin/golang-starterkit/packages/post"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter registers all application routes
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Register global middlewares
	r.Use(middleware.CORS())

	// Initialize handlers
	postRepo := post.NewRepository(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)

	// API version 1 group
	v1 := r.Group("/api/v1")
	{
		// Blog/Post endpoints
		posts := v1.Group("/posts")
		{
			// Public routes
			posts.GET("", postHandler.GetPosts)
			posts.GET("/:id", postHandler.GetPost)

			// Protected routes (Requires Auth)
			protected := posts.Group("")
			protected.Use(middleware.AuthMiddleware(cfg))
			{
				// Only admins can create, update, or delete posts
				protected.POST("", middleware.AuthorizeRoles("admin"), postHandler.CreatePost)
				protected.PUT("/:id", middleware.AuthorizeRoles("admin"), postHandler.UpdatePost)
				protected.DELETE("/:id", middleware.AuthorizeRoles("admin"), postHandler.DeletePost)
			}
		}
	}

	return r
}
