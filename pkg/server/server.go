package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/matherique/middleware/internal/app"
)

type server struct {
	a   app.APP
	ctx context.Context
	log *log.Logger
}

func NewServer(ctx context.Context, a app.APP, log *log.Logger) *server {
	return &server{
		a:   a,
		ctx: ctx,
		log: log,
	}
}

func (s *server) Listen(port string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", s.GetUsersHandlerFunc)

	return s.start(port, mux)
}

func (s *server) start(port string, mux *http.ServeMux) error {
	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			s.log.Fatal(err)
		}
	}()

	s.log.Printf("server start in: localhost%s\n", port)

	<-s.ctx.Done()

	s.log.Println("server stoped")

	ctxSD, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctxSD); err != nil {
		return err
	}

	s.log.Println("server exited properly")

	return nil
}

func (s *server) GetUsersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var data app.GetUsersRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("invalid body"))
		return
	}

	resp, err := s.a.GetUsers(s.ctx, data)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(resp)
}
