package chi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	srv *http.Server
}

func Start(port string) *Server {
	r := chi.NewRouter()

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, World!",
			"framework": "Chi",
		})
	})

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	server := &Server{srv: srv}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Chi server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.Background())
}

func Main() {
	log.Println("Chi server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
