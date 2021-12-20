package middleware

import (
	"context"
	"log"
	"time"

	"github.com/matherique/middleware/internal/api"
)

type middleware struct {
	api api.API
	log *log.Logger
}

func NewMiddleware(a api.API, logger *log.Logger) api.API {
	return &middleware{a, logger}
}

func (m *middleware) GetUsers(ctx context.Context, req api.GetUsersRequest) (*api.GetUsersResponse, error) {
	defer func(start time.Time) {
		finish := time.Since(start)

		m.log.SetPrefix("[GetUsers] ")
		m.log.Printf("finish in %s", finish)
	}(time.Now())

	return m.api.GetUsers(ctx, req)
}

func (m *middleware) GetBooks(ctx context.Context, req api.GetBooksRequest) (*api.GetBooksResponse, error) {
	defer func(start time.Time) {
		finish := time.Since(start)

		m.log.SetPrefix("[GetBooks] ")
		m.log.Printf("finish in %s", finish)
	}(time.Now())

	return m.api.GetBooks(ctx, req)
}
