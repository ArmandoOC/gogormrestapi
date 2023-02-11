package models

import "gorm.io/gorm"

// Article structure for our blog
type Article struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Content     string `gorm:"not null" json:"content"`
	Description string `json:"description"`
}
