package repository

import (
	"techdocs/internal/model"
	"techdocs/pkg/logger"

	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

// Create creates a new document in the database
func (r *DocumentRepository) Create(document *model.Document) error {
	return r.db.Create(document).Error
}

// Update updates an existing document in the database
func (r *DocumentRepository) Update(document *model.Document) error {
	return r.db.Save(document).Error
}

// Delete deletes a document from the database
func (r *DocumentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Document{}, id).Error
}

// GetByID retrieves a document by its ID
func (r *DocumentRepository) GetByID(id uint) (*model.Document, error) {
	logger.Info("Repository: Getting document by ID: %d", id)
	var document model.Document
	err := r.db.Preload("Author").Preload("Tags").Preload("Service").First(&document, id).Error
	if err != nil {
		logger.Error("Repository: Error getting document by ID %d: %v", id, err)
		return nil, err
	}
	logger.Info("Repository: Successfully retrieved document: %d", id)
	return &document, nil
}

// GetAll retrieves all documents
func (r *DocumentRepository) GetAll() ([]model.Document, error) {
	var documents []model.Document
	err := r.db.Preload("Author").Preload("Tags").Preload("Service").Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

// GetByAuthor retrieves all documents by an author ID
func (r *DocumentRepository) GetByAuthor(authorID uint) ([]model.Document, error) {
	var documents []model.Document
	err := r.db.Where("author_id = ?", authorID).Preload("Author").Preload("Tags").Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

// GetByCategory retrieves all documents by category
func (r *DocumentRepository) GetByCategory(category string) ([]model.Document, error) {
	var documents []model.Document
	err := r.db.Where("category = ?", category).Preload("Author").Preload("Tags").Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (r *DocumentRepository) GetByAuthorID(authorID uint) ([]model.Document, error) {
	var documents []model.Document
	err := r.db.Where("author_id = ?", authorID).Preload("Author").Preload("Tags").Preload("Service").Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}
