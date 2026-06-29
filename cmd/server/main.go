package main

import (
	"context"
	"log"

	"github.com/MHG14/aethoria_marketplace/internal/application"
	"github.com/MHG14/aethoria_marketplace/internal/domain/ports/service"
	"github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres"
	"github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/repository"
	httpserver "github.com/MHG14/aethoria_marketplace/internal/transport/http"
	"github.com/MHG14/aethoria_marketplace/internal/transport/http/handlers"
)

func main() {
	ctx := context.Background()

	pool, err := postgres.NewPool(ctx, "postgres://myuser:mypassword@localhost:5499/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repos := repository.New(pool)
	app := application.NewApp(repos, service.Services{})
	h := handlers.New(app)
	srv := httpserver.NewServer(h)

	log.Fatal(srv.Listen(":8080"))
}
