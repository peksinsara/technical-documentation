package handler

import (
	"net/http"
	"strconv"
	"techdocs/internal/model"
	"techdocs/internal/service"
	"techdocs/pkg/logger"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	documentService *service.DocumentService
}

func NewDocumentHandler(documentService *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		documentService: documentService,
	}
}

// RegisterRoutes registers the document routes
func (h *DocumentHandler) RegisterRoutes(router *gin.RouterGroup) {
	documents := router.Group("/documents")
	{
		documents.POST("", h.CreateDocument)
		documents.PUT("/:id", h.UpdateDocument)
		documents.DELETE("/:id", h.DeleteDocument)
		documents.GET("/:id", h.GetDocumentByID)
		documents.GET("", h.GetAllDocuments)
		documents.GET("/author/:authorID", h.GetDocumentsByAuthor)
		documents.GET("/category/:category", h.GetDocumentsByCategory)
	}
}

// CreateDocument handles the creation of a new document
func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var doc model.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		logger.Error("Document creation validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Set the author ID
	doc.AuthorID = userID.(uint)

	if err := h.documentService.CreateDocument(&doc); err != nil {
		logger.Error("Failed to create document for user ID %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Document created successfully: ID %d by user ID %d", doc.ID, userID)
	c.JSON(http.StatusCreated, doc)
}

// UpdateDocument handles the update of an existing document
func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	var doc model.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		logger.Error("Document update validation error for ID %s: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")
	doc.AuthorID = userID

	if err := h.documentService.UpdateDocument(&doc); err != nil {
		logger.Error("Failed to update document ID %s for user ID %d: %v", id, userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Document updated successfully: ID %s by user ID %d", id, userID)
	c.JSON(http.StatusOK, doc)
}

// DeleteDocument handles the deletion of a document
func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Invalid document ID format: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID format"})
		return
	}

	if err := h.documentService.DeleteDocument(uint(id)); err != nil {
		logger.Error("Failed to delete document ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Document deleted successfully: ID %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

// GetDocumentByID handles the retrieval of a document by its ID
func (h *DocumentHandler) GetDocumentByID(c *gin.Context) {
	idStr := c.Param("id")
	logger.Info("Received document ID parameter: %s", idStr)
	logger.Info("Request URL: %s", c.Request.URL.String())
	logger.Info("Request method: %s", c.Request.Method)
	logger.Info("Request path: %s", c.Request.URL.Path)

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Error("Invalid document ID format: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID format"})
		return
	}

	logger.Info("Parsed document ID: %d", id)

	doc, err := h.documentService.GetDocumentByID(uint(id))
	if err != nil {
		logger.Error("Failed to get document ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	logger.Info("Successfully retrieved document: %d", id)
	c.JSON(http.StatusOK, doc)
}

// GetAllDocuments handles the retrieval of all documents
func (h *DocumentHandler) GetAllDocuments(c *gin.Context) {
	docs, err := h.documentService.GetAllDocuments()
	if err != nil {
		logger.Error("Failed to get all documents: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, docs)
}

// GetDocumentsByAuthor handles the retrieval of documents by author ID
func (h *DocumentHandler) GetDocumentsByAuthor(c *gin.Context) {
	authorIDStr := c.Param("authorID")
	authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
	if err != nil {
		logger.Error("Invalid author ID format: %s", authorIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID format"})
		return
	}

	docs, err := h.documentService.GetDocumentsByAuthor(uint(authorID))
	if err != nil {
		logger.Error("Failed to get documents for author ID %d: %v", authorID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, docs)
}

// GetDocumentsByCategory handles the retrieval of documents by category
func (h *DocumentHandler) GetDocumentsByCategory(c *gin.Context) {
	category := c.Param("category")
	docs, err := h.documentService.GetDocumentsByCategory(category)
	if err != nil {
		logger.Error("Failed to get documents for category %s: %v", category, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, docs)
}
