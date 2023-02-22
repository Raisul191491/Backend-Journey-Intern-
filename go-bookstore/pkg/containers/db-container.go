package containers

import (
	"github.com/deadking/go-bookstore/pkg/config"
	"github.com/deadking/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDatabse() *gorm.DB {
	config.Connect()
	db = config.GetDB()

	db.Migrator().AutoMigrate(&models.Author{})
	db.Migrator().AutoMigrate(&models.Book{})
	// db.Migrator().DropTable(&models.Book{}, &models.Author{})
	// db.Migrator().CreateTable(&models.Author{})
	// db.Migrator().CreateTable(&models.Book{})

	return db
}
