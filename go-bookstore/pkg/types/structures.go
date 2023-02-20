package types

type CustomBookResponse struct {
	Content *ResponseBook `json:"content,omitempty"`
	Msg     string
}
type CustomAuthorResponse struct {
	Content *ResponseAuthor `json:"content,omitempty"`
	Msg     string
}

type CustomDeleteResponse struct {
	Msg string
}

type ResponseBook struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Publication string `json:"publication,omitempty"`
	AuthorID    uint   `json:"author_id,omitempty"`
	Author      ResponseAuthor
}

type ResponseAuthor struct {
	ID         uint   `json:"id,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Age        int    `json:"age,omitempty"`
}
