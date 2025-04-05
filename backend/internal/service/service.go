package service

import (
	"techdocs/internal/model"
	"techdocs/internal/repository"
)

type ServiceService struct {
	repo *repository.ServiceRepository
}

func NewServiceService(repo *repository.ServiceRepository) *ServiceService {
	return &ServiceService{repo: repo}
}

// CreateService creates a new service
func (s *ServiceService) CreateService(service *model.Service) error {
	return s.repo.Create(service)
}

// UpdateService updates an existing service
func (s *ServiceService) UpdateService(service *model.Service) error {
	return s.repo.Update(service)
}

// DeleteService deletes a service
func (s *ServiceService) DeleteService(id uint) error {
	return s.repo.Delete(id)
}

// GetServiceByID retrieves a service by its ID
func (s *ServiceService) GetServiceByID(id uint) (*model.Service, error) {
	return s.repo.GetByID(id)
}

// GetAllServices retrieves all services
func (s *ServiceService) GetAllServices() ([]model.Service, error) {
	return s.repo.GetAll()
}

// GetServicesByCategory retrieves all services by category
func (s *ServiceService) GetServicesByCategory(category string) ([]model.Service, error) {
	return s.repo.GetByCategory(category)
}
