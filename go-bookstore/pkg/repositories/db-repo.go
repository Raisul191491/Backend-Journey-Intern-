package repositories

import "gorm.io/gorm"

var db *gorm.DB

func DbInstance(d *gorm.DB) {
	db = d
}
