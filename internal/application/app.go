package application

import (
	"github.com/MHG14/aethoria_marketplace/internal/domain/ports/repository"
	"github.com/MHG14/aethoria_marketplace/internal/domain/ports/service"
)

// App is the single entry point for all use cases.
// Handlers only talk to this — never to repos or domain directly.
type App struct {
	repos repository.Repositories
	svc   service.Services
}

func NewApp(repos repository.Repositories, svc service.Services) *App {
	return &App{repos: repos, svc: svc}
}
