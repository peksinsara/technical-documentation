package main

import (
	"log"
	"techdocs/internal/config"
	"techdocs/internal/handler"
	"techdocs/internal/middleware"
	"techdocs/internal/repository"
	"techdocs/internal/service"
	"techdocs/pkg/database"
	"techdocs/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	if err := logger.Init("./logs"); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	logger.Info("Application starting...")

	// Initialize database
	db, err = database.NewMySQLDB(cfg)
	if err != nil {
		logger.Error("Failed to initialize database: %v", err)
		log.Fatalf("Failed to initialize database: %v", err)
	}
	logger.Info("Database initialized successfully")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	documentRepo := repository.NewDocumentRepository(db)
	serviceRepo := repository.NewServiceRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	documentService := service.NewDocumentService(documentRepo)
	serviceService := service.NewServiceService(serviceRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	documentHandler := handler.NewDocumentHandler(documentService)
	serviceHandler := handler.NewServiceHandler(serviceService)

	// Initialize Gin router
	router := gin.Default()

	// Add logger middleware
	router.Use(middleware.LoggerMiddleware())

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Setup routes
	userHandler.SetupRoutes(router)

	// API routes group
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		documentHandler.RegisterRoutes(api)
		serviceHandler.RegisterRoutes(api)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	addr := ":" + cfg.ServerPort
	logger.Info("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		logger.Error("Failed to start server: %v", err)
		log.Fatalf("Failed to start server: %v", err)
	}
}
