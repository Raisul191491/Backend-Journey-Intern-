package models

type Author struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true"`
	AuthorName string `json:"author_name"`
	Age        int    `json:"age"`
}

func (a Author) CreateAuthor() (*Author, string) {
	return &a, "Hi"
}
