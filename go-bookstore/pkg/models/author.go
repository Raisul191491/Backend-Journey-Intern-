package models

type Author struct {
	ID         uint `gorm:"primaryKey"`
	AuthorName string
	Age        int
}