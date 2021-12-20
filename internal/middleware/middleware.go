package middleware

import (
	"context"
	"log"
	"time"

	"github.com/matherique/middleware/internal/app"
)

type middleware struct {
	app app.APP
	log *log.Logger
}

func NewMiddleware(a app.APP, logger *log.Logger) app.APP {
	return &middleware{a, logger}
}

func (m *middleware) GetUsers(ctx context.Context, req app.GetUsersRequest) (*app.GetUsersResponse, error) {
	defer func(start time.Time) {
		finish := time.Since(start)

		m.log.SetPrefix("[GetUsers] ")
		m.log.Printf("finish in %s", finish)
	}(time.Now())

	return m.app.GetUsers(ctx, req)
}

func (m *middleware) GetBooks(ctx context.Context, req app.GetBooksRequest) (*app.GetBooksResponse, error) {
	defer func(start time.Time) {
		finish := time.Since(start)

		m.log.SetPrefix("[GetBooks] ")
		m.log.Printf("finish in %s", finish)
	}(time.Now())

	return m.app.GetBooks(ctx, req)
}
