package fiber

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func Start(port string) *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
			"framework": "Fiber",
		})
	})

	server := &Server{app: app}

	go func() {
		if err := app.Listen(port); err != nil {
			log.Printf("Fiber server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	return s.app.Shutdown()
}

func Main() {
	log.Println("Fiber server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
