package database

import (
	"fmt"
	"techdocs/internal/config"
	"techdocs/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(
		&model.User{},
		&model.Document{},
		&model.Tag{},
		&model.Comment{},
		&model.DocumentVersion{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
} 