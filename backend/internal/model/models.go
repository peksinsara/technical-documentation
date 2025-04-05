package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	Role      string     `gorm:"type:varchar(50);not null;default:'user'" json:"role"`
	Documents []Document `gorm:"foreignKey:AuthorID" json:"documents,omitempty"`
}

type Document struct {
	gorm.Model
	Title       string   `json:"title" gorm:"not null"`
	Description string   `json:"description"`
	Content     string   `json:"content" gorm:"type:text"`
	Type        string   `json:"type" gorm:"not null"`
	Category    string   `json:"category"`
	AuthorID    uint     `json:"author_id" gorm:"not null"`
	Author      User     `json:"author" gorm:"foreignKey:AuthorID"`
	Tags        []Tag    `json:"tags" gorm:"many2many:document_tags;"`
	ServiceID   *uint    `json:"service_id"`
	Service     *Service `json:"service"`
}

type Tag struct {
	gorm.Model
	Name      string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
	Documents []Document `gorm:"many2many:document_tags;" json:"documents,omitempty"`
}

type Comment struct {
	gorm.Model
	Content    string   `gorm:"type:text;not null" json:"content"`
	DocumentID uint     `gorm:"not null" json:"document_id"`
	Document   Document `gorm:"foreignKey:DocumentID" json:"document,omitempty"`
	UserID     uint     `gorm:"not null" json:"user_id"`
	User       User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type DocumentVersion struct {
	gorm.Model
	DocumentID uint      `gorm:"not null" json:"document_id"`
	Document   Document  `gorm:"foreignKey:DocumentID" json:"document,omitempty"`
	Content    string    `gorm:"type:text" json:"content"`
	Version    int       `gorm:"not null" json:"version"`
	CreatedBy  uint      `gorm:"not null" json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type Service struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}
