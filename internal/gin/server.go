package gin

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
}

func Start(port string) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "Hello, World!",
			"framework": "Gin",
		})
	})

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	server := &Server{srv: srv}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Gin server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.Background())
}

func Main() {
	log.Println("Gin server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
