package app

import "context"

type APP interface {
	GetUsers(context.Context, GetUsersRequest) (*GetUsersResponse, error)
	GetBooks(context.Context, GetBooksRequest) (*GetBooksResponse, error)
}

type app struct {
}

func NewApi() APP {
	return new(app)
}

type GetUsersRequest struct {
	Id int `json:"id"`
}

type GetUsersResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type GetBooksRequest struct {
	CategoryId int `json:"category_id"`
	AuthorId   int `json:"author_id"`
}

type GetBooksResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	NumberPages int    `json:"number_pages"`
}
