package containers

import (
	"github.com/deadking/go-bookstore/pkg/repositories"
	"github.com/deadking/go-bookstore/pkg/services"
	"gorm.io/gorm"
)

func Initializeinterfaces(db *gorm.DB) {
	Ibookcrud := repositories.BookDbInstance(db)
	Iauthorcrud := repositories.AuthorDbInstance(db)
	services.BookInterfaceInstance(Ibookcrud)
	services.AuthorInterfaceInstance(Iauthorcrud)
}
