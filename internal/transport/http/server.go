package httpserver

import (
	"github.com/MHG14/aethoria_marketplace/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server struct {
	app      *fiber.App
	handlers *handlers.Handlers
}

func NewServer(h *handlers.Handlers) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${method} ${path} ${latency}\n",
	}))

	s := &Server{app: app, handlers: h}
	s.registerRoutes()
	return s
}

func (s *Server) Listen(addr string) error {
	return s.app.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
