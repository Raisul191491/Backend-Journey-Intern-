package types

type CustomBookResponse struct {
	Content  *ResponseBook `json:"Book,omitempty"`
	ErrorMsg error         `json:"Error message,omitempty"`
	Msg      string
}
type CustomAuthorResponse struct {
	Content  *ResponseAuthor `json:"Author,omitempty"`
	ErrorMsg error           `json:"Error message,omitempty"`
	Msg      string
}

type CustomOnlyResponse struct {
	Msg string
}

type ResponseBook struct {
	ID          uint           `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Publication string         `json:"publication,omitempty"`
	AuthorID    uint           `json:"author_id,omitempty"`
	Author      ResponseAuthor `json:"author,omitempty"`
}

type ResponseAuthor struct {
	ID         uint   `json:"id,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Age        int    `json:"age,omitempty"`
}
