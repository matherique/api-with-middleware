package app

import (
	"context"
	"time"
)

type Users interface {
	GetUsers(context.Context, GetUsersRequest) (*GetUsersResponse, error)
}

func (a *app) GetUsers(ctx context.Context, req GetUsersRequest) (*GetUsersResponse, error) {
	time.Sleep(5 * time.Second)

	resp := &GetUsersResponse{
		Name:  "Matheus Henrique",
		Age:   27,
		Email: "matherique@gmail.com",
	}

	return resp, nil
}

type GetUsersRequest struct {
	Id int `json:"id"`
}

type GetUsersResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
