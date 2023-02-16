package repositories

import "github.com/deadking/go-bookstore/pkg/models"

// type TempBook struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Publication string `json:"publication"`
// 	AuthorID    uint   `json:"author_id"`
// }

func CreateBook(b models.Book) (models.Book, string) {
	err := b.Validate()
	if err == nil {
		db.Table("books").Create(&b)
		return b, "Book created, Successfully"
	}
	return b, err.Error()
}

func DeleteBook(ID int) (models.Book, string) {
	var deletedBook models.Book

	db.Where("ID=?", ID).Find(&deletedBook)
	if deletedBook.Name == "" || deletedBook.Publication == "" {
		return deletedBook, "Book not found to begin with"
	}
	db.Where("ID=?", ID).Delete(&deletedBook)
	return deletedBook, "Successfully deleted...."
}

func UpdateBook(ID int, updateBook models.Book) (models.Book, string) {
	var book models.Book
	db.Where("ID=?", ID).Find(&book)
	if book.Name == "" || book.Publication == "" {
		return book, "Book not found"
	}

	// Update or reject update
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	err := book.Validate()
	if err == nil {
		db.Save(&book)
		return book, "Successfully updated"
	}
	return book, err.Error()
}

func GetBook(bookID, authorID int) []models.Book {
	var books []models.Book

	if bookID > 0 {
		db.Where("ID=?", bookID).Find(&books)
	}
	if authorID > 0 {
		db.Where("author_id=?", authorID).Find(&books)
	}
	if bookID == 0 && authorID == 0 {
		db.Find(&books)
	}
	return books
}
