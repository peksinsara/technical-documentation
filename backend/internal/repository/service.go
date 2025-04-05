package repository

import (
	"techdocs/internal/model"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

// Create creates a new service in the database
func (r *ServiceRepository) Create(service *model.Service) error {
	return r.db.Create(service).Error
}

// Update updates an existing service in the database
func (r *ServiceRepository) Update(service *model.Service) error {
	return r.db.Save(service).Error
}

// Delete deletes a service from the database
func (r *ServiceRepository) Delete(id uint) error {
	return r.db.Delete(&model.Service{}, id).Error
}

// GetByID retrieves a service by its ID
func (r *ServiceRepository) GetByID(id uint) (*model.Service, error) {
	var service model.Service
	err := r.db.First(&service, id).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

// GetAll retrieves all services
func (r *ServiceRepository) GetAll() ([]model.Service, error) {
	var services []model.Service
	err := r.db.Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

// GetByCategory retrieves all services by category
func (r *ServiceRepository) GetByCategory(category string) ([]model.Service, error) {
	var services []model.Service
	err := r.db.Where("category = ?", category).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}
