package handler

import (
	"net/http"
	"techdocs/internal/middleware"
	"techdocs/internal/model"
	"techdocs/internal/service"
	"techdocs/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Registration validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Register(&req); err != nil {
		logger.Error("Registration failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Info("User registered successfully: %s", req.Username)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Login validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userService.Login(&req)
	if err != nil {
		logger.Error("Login failed for user %s: %v", req.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	logger.Info("User logged in successfully: %s", req.Username)
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		logger.Error("Failed to get profile for user ID %d: %v", userID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	var updates model.User
	if err := c.ShouldBindJSON(&updates); err != nil {
		logger.Error("Profile update validation error for user ID %d: %v", userID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.UpdateUser(userID, &updates); err != nil {
		logger.Error("Failed to update profile for user ID %d: %v", userID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Profile updated successfully for user ID %d", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	page := 1
	limit := 10
	// TODO: Get page and limit from query parameters

	users, total, err := h.userService.ListUsers(page, limit)
	if err != nil {
		logger.Error("Failed to list users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": total,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.GetUint("userID")
	if err := h.userService.DeleteUser(userID); err != nil {
		logger.Error("Failed to delete user ID %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("User deleted successfully: ID %d", userID)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/api/auth/register", h.Register)
	router.POST("/api/auth/login", h.Login)

	// Protected routes
	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware(h.userService.GetJWTSecret()))
	{
		auth.GET("/profile", h.GetProfile)
		auth.PUT("/profile", h.UpdateProfile)
		auth.DELETE("/profile", h.DeleteUser)

		// Admin routes
		admin := auth.Group("/admin")
		admin.Use(middleware.RequireRole("admin"))
		{
			admin.GET("/users", h.ListUsers)
		}
	}
}
