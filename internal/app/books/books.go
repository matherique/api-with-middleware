package app

import (
	"context"
	"time"
)

type Books interface {
	GetBooks(context.Context, GetBooksRequest) (*GetBooksResponse, error)
}

func (a *app) GetBooks(ctx context.Context, req GetBooksRequest) (*GetBooksResponse, error) {
	time.Sleep(2 * time.Second)

	resp := &GetBooksResponse{
		Id:          1,
		Name:        "Harry Potter",
		Description: "Harry potter é um bruxo chato",
		NumberPages: 400,
	}

	return resp, nil
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
