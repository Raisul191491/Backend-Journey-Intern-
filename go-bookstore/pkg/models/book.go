package models

// var db *gorm.DB

type Book struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Publication string
	AuthorID    uint
	Author      Author `gorm:"foreignKey:AuthorID"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&b.Author, validation.Required, validation.Length(5, 50)),
	)
}

func (b *Book) CreateBook() (*Book, string) {
	err := b.Validate()
	if err == nil {
		db.NewRecord(b)
		db.Create(&b)
		success := "Successfully updated"
		return b, success
	}
	return b, err.Error()
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

// func GetBookAnyway(Id int64) ([]Book, *gorm.DB) {
// 	var books []Book
// 	if Id > 0 {
// 		db.Where("ID=?", Id).Find(&books)
// 	} else {
// 		db.Find(&books)
// 	}
// 	return books, db
// }

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
