package models

// Book ...
type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Chapters []*Chapter `orm:"reverse(many)" json:"-"`
}

type RespDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BooksDTO struct {
	RespDTO
	Data []*Book `json:"data"`
}

type BookDTO struct {
	RespDTO
	Data Book `json:"data"`
}
