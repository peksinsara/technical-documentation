package service

import (
	"techdocs/internal/model"
	"techdocs/internal/repository"
)

type DocumentService struct {
	repo *repository.DocumentRepository
}

func NewDocumentService(repo *repository.DocumentRepository) *DocumentService {
	return &DocumentService{repo: repo}
}

// CreateDocument creates a new document
func (s *DocumentService) CreateDocument(document *model.Document) error {
	return s.repo.Create(document)
}

// UpdateDocument updates an existing document
func (s *DocumentService) UpdateDocument(document *model.Document) error {
	return s.repo.Update(document)
}

// DeleteDocument deletes a document
func (s *DocumentService) DeleteDocument(id uint) error {
	return s.repo.Delete(id)
}

// GetDocumentByID retrieves a document by its ID
func (s *DocumentService) GetDocumentByID(id uint) (*model.Document, error) {
	return s.repo.GetByID(id)
}

// GetAllDocuments retrieves all documents
func (s *DocumentService) GetAllDocuments() ([]model.Document, error) {
	return s.repo.GetAll()
}

// GetDocumentsByAuthor retrieves all documents by an author ID
func (s *DocumentService) GetDocumentsByAuthor(authorID uint) ([]model.Document, error) {
	return s.repo.GetByAuthor(authorID)
}

// GetDocumentsByCategory retrieves all documents by category
func (s *DocumentService) GetDocumentsByCategory(category string) ([]model.Document, error) {
	return s.repo.GetByCategory(category)
}
