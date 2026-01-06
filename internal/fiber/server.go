package fiber

import (
	"context"
	"log"

	"github.com/creativeyann17/GolangAPIsTest/database"
	"github.com/creativeyann17/GolangAPIsTest/repository"
	"github.com/creativeyann17/GolangAPIsTest/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const dbConnString = "postgres://postgres:postgres@localhost:5432/golang_apis?sslmode=disable"

type Server struct {
	app *fiber.App
	db  *database.PostgresDB
}

func Start(port string) *Server {
	ctx := context.Background()

	db, err := database.NewPostgresDB(ctx, dbConnString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	cityRepo := repository.NewCityRepository(db.Pool)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "Hello, World!",
			"framework": "Fiber",
		})
	})

	// City CRUD routes
	app.Post("/cities", func(c *fiber.Ctx) error {
		var city types.City
		if err := c.BodyParser(&city); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		dao := city.ToCityDAO()

		if err := cityRepo.Create(c.Context(), &dao); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(dao.ToCity())
	})

	app.Get("/cities", func(c *fiber.Ctx) error {
		daos, err := cityRepo.GetAll(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		cities := make([]types.City, len(daos))
		for i, dao := range daos {
			cities[i] = dao.ToCity()
		}

		return c.JSON(cities)
	})

	app.Get("/cities/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		dao, err := cityRepo.GetByID(c.Context(), id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if dao == nil {
			return c.Status(404).JSON(fiber.Map{"error": "City not found"})
		}

		return c.JSON(dao.ToCity())
	})

	app.Put("/cities/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		var city types.City
		if err := c.BodyParser(&city); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		dao := city.ToCityDAO()
		dao.ID = id

		if err := cityRepo.Update(c.Context(), &dao); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(dao.ToCity())
	})

	app.Delete("/cities/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		if err := cityRepo.Delete(c.Context(), id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(204)
	})

	server := &Server{app: app, db: db}

	go func() {
		if err := app.Listen(port); err != nil {
			log.Printf("Fiber server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	if s.db != nil {
		s.db.Close()
	}
	return s.app.Shutdown()
}

func Main() {
	log.Println("Fiber server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
