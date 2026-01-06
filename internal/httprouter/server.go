package httprouter

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	srv *http.Server
}

func Start(port string) *Server {
	router := httprouter.New()

	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, World!",
			"framework": "HttpRouter",
		})
	})

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	server := &Server{srv: srv}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HttpRouter server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.Background())
}

func Main() {
	log.Println("HttpRouter server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
