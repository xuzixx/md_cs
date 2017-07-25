package models

// RespDTO http return
type RespDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BooksDTO http return
type BooksDTO struct {
	RespDTO
	Data []*Book `json:"data"`
}

// BookDTO http return
type BookDTO struct {
	RespDTO
	Data Book `json:"data"`
}
