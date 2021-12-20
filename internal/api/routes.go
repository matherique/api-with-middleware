package api

import (
	"context"
	"time"
)

func (a *api) GetUsers(ctx context.Context, req GetUsersRequest) (*GetUsersResponse, error) {
	time.Sleep(5 * time.Second)

	resp := &GetUsersResponse{
		Name:  "Matheus Henrique",
		Age:   27,
		Email: "matherique@gmail.com",
	}

	return resp, nil
}

func (a *api) GetBooks(ctx context.Context, req GetBooksRequest) (*GetBooksResponse, error) {
	time.Sleep(2 * time.Second)

	resp := &GetBooksResponse{
		Id:          1,
		Name:        "Harry Potter",
		Description: "Harry potter é um bruxo chato",
		NumberPages: 400,
	}

	return resp, nil
}
