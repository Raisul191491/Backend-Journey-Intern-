package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	Publication string `json:"publication"`
	AuthorID    uint   `json:"-"`
	Author      Author `json:"author" gorm:"foreignKey:AuthorID ;references:ID"`
	// AuthorID    uint   `json:"-"`
	// Author      Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type TempBook struct {
	Name        string `json:"name"`
	Publication string `json:"publication"`
	AuthorID    uint   `json:"author_id"`
}

type Author struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	AuthorName string `json:"author_name"`
	Age        int    `json:"age"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&b.Publication, validation.Required, validation.Length(6, 50)),
	)
}

func (b TempBook) CreateBook() (*TempBook, string) {
	//err := b.Validate()

	// fmt.Println("sdfs", b)
	// if err == nil {
	// db.NewRecord(b)

	temp := Book{
		Name:        b.Name,
		Publication: b.Publication,
		AuthorID:    b.AuthorID,
	}

	fmt.Println(b, "last")

	db.Joins("Author").Create(&temp)
	success := "Successfully updated"
	return &b, success
	// }
	// if err != nil {
	// 	return b, "Pain is everywhere"
	// }
	// return b, "Kill me now"
}

// func (b *Book) UpdateBook() (*Book, string) {
// 	err := b.Validate()
// 	if err == nil {
// 		db.Save(&b)
// 		success := "Successfully updated"
// 		return b, success
// 	}
// 	return b, err.Error()
// }

func GetBookAnyway(Id int64) ([]Book, *gorm.DB) {
	var books []Book
	if Id > 0 {
		db.Where("ID=?", Id).Find(&books)
	} else {
		fmt.Println(db.Find(&books).Error)
		db.Find(&books)
	}
	return books, db
}

// func DeleteBook(Id int64) (Book, string) {
// 	var book, deletedBook Book
// 	db.Where("ID=?", Id).Find(&deletedBook)
// 	if deletedBook.Name == "" || deletedBook.Author == "" {
// 		return deletedBook, "Book not found to begin with"
// 	}
// 	db.Where("ID=?", Id).Delete(&book)
// 	success := "Successfully deleted...."
// 	return deletedBook, success
// }
