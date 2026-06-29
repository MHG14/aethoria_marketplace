package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/guild"
)

type GetGuildResponse struct {
	Guild guild.Guild `json:"guild"`
}

func (a *App) GetGuild(ctx context.Context, id int64) (GetGuildResponse, error) {
	g, err := a.repos.Guild.GetByID(ctx, id)
	if err != nil {
		return GetGuildResponse{}, fmt.Errorf("get guild: %w", err)
	}
	return GetGuildResponse{Guild: g}, nil
}
