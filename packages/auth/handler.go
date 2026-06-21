package auth

import (
	"net/http"

	"github.com/alirfanyasin/golang-starterkit/config"
	"github.com/alirfanyasin/golang-starterkit/middleware"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"password"`
	Role     string `json:"role" example:"admin"` // admin, member
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

// Login godoc
// @Summary      Authenticate User & Get JWT Token
// @Description  Simple auth mock endpoint to get a JWT token. Enter "admin" and "password" with "admin" role.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials body      LoginInput  true  "Login Credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simple mock validation (any user with password "password" is allowed)
	if input.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	role := input.Role
	if role == "" {
		role = "admin" // Default to admin for easier testing
	}

	// Generate JWT token (user ID is mocked as 1)
	token, err := middleware.GenerateToken(1, role, h.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
