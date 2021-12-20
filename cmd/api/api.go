package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/matherique/middleware/internal/app"
	"github.com/matherique/middleware/internal/middleware"
	"github.com/matherique/middleware/pkg/server"
)

func main() {
	a := app.NewApi()
	log := log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lmsgprefix)
	a = middleware.NewMiddleware(a, log)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	srv := server.NewServer(ctx, a, log)

	go func() {
		call := <-c
		log.Printf("system call: %+v", call)

		cancel()
	}()

	if err := srv.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
