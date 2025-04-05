package handler

import (
	"net/http"
	"strconv"
	"techdocs/internal/model"
	"techdocs/internal/service"
	"techdocs/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	serviceService *service.ServiceService
}

func NewServiceHandler(serviceService *service.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		serviceService: serviceService,
	}
}

// RegisterRoutes registers the service routes
func (h *ServiceHandler) RegisterRoutes(router *gin.RouterGroup) {
	services := router.Group("/services")
	{
		services.POST("", h.CreateService)
		services.PUT("/:id", h.UpdateService)
		services.DELETE("/:id", h.DeleteService)
		services.GET("/:id", h.GetServiceByID)
		services.GET("", h.GetAllServices)
		services.GET("/category/:category", h.GetServicesByCategory)
	}
}

// CreateService handles the creation of a new service
func (h *ServiceHandler) CreateService(c *gin.Context) {
	var svc model.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		logger.Error("Service creation validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.serviceService.CreateService(&svc); err != nil {
		logger.Error("Failed to create service: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Service created successfully: ID %d", svc.ID)
	c.JSON(http.StatusCreated, svc)
}

// UpdateService handles the update of an existing service
func (h *ServiceHandler) UpdateService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Invalid service ID format: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID format"})
		return
	}

	var svc model.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		logger.Error("Service update validation error for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	svc.ID = uint(id)

	if err := h.serviceService.UpdateService(&svc); err != nil {
		logger.Error("Failed to update service ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Service updated successfully: ID %d", id)
	c.JSON(http.StatusOK, svc)
}

// DeleteService handles the deletion of a service
func (h *ServiceHandler) DeleteService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Invalid service ID format: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID format"})
		return
	}

	if err := h.serviceService.DeleteService(uint(id)); err != nil {
		logger.Error("Failed to delete service ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Service deleted successfully: ID %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

// GetServiceByID handles the retrieval of a service by its ID
func (h *ServiceHandler) GetServiceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Invalid service ID format: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID format"})
		return
	}

	svc, err := h.serviceService.GetServiceByID(uint(id))
	if err != nil {
		logger.Error("Failed to get service ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, svc)
}

// GetAllServices handles the retrieval of all services
func (h *ServiceHandler) GetAllServices(c *gin.Context) {
	services, err := h.serviceService.GetAllServices()
	if err != nil {
		logger.Error("Failed to get all services: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}

// GetServicesByCategory handles the retrieval of services by category
func (h *ServiceHandler) GetServicesByCategory(c *gin.Context) {
	category := c.Param("category")
	services, err := h.serviceService.GetServicesByCategory(category)
	if err != nil {
		logger.Error("Failed to get services for category %s: %v", category, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}
