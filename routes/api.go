package routes

import (
	_ "github.com/alirfanyasin/golang-starterkit/docs"
	"github.com/alirfanyasin/golang-starterkit/config"
	"github.com/alirfanyasin/golang-starterkit/middleware"
	"github.com/alirfanyasin/golang-starterkit/packages/auth"
	"github.com/alirfanyasin/golang-starterkit/packages/post"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// SetupRouter registers all application routes
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register global middlewares
	r.Use(middleware.CORS())

	// Initialize handlers
	postRepo := post.NewRepository(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)
	authHandler := auth.NewHandler(cfg)

	// API version 1 group
	v1 := r.Group("/api/v1")
	{
		// Authentication endpoints
		v1.POST("/auth/login", authHandler.Login)
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
