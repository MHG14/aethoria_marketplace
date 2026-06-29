package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/guild"
)

type CreateGuildRequest struct {
	Name       string `json:"name"`
	DailyLimit int64  `json:"daily_limit"`
}

type CreateGuildResponse struct {
	Guild guild.Guild `json:"guild"`
}

func (a *App) CreateGuild(ctx context.Context, req CreateGuildRequest) (CreateGuildResponse, error) {
	if req.Name == "" {
		return CreateGuildResponse{}, fmt.Errorf("%w: name is required", ErrInvalidInput)
	}
	if req.DailyLimit <= 0 {
		return CreateGuildResponse{}, fmt.Errorf("%w: daily_limit must be positive", ErrInvalidInput)
	}

	g, err := a.repos.Guild.Create(ctx, guild.Guild{
		Name:       req.Name,
		DailyLimit: req.DailyLimit,
	})
	if err != nil {
		return CreateGuildResponse{}, fmt.Errorf("create guild: %w", err)
	}

	return CreateGuildResponse{Guild: g}, nil
}
