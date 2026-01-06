package echo

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func Start(port string) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
			"framework": "Echo",
		})
	})

	server := &Server{echo: e}

	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			log.Printf("Echo server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	return s.echo.Shutdown(context.Background())
}

func Main() {
	log.Println("Echo server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
