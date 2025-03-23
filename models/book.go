package models

import "gorm.io/gorm"

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:255;not null"`
	Author string `gorm:"size:255;not null"`
	Year   int    `gorm:"not null"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
