package application

import (
	"context"
	"fmt"
)

type GetWalletResponse struct {
	GuildID          int64 `json:"guild_id"`
	TotalMoney       int64 `json:"total_money"`
	ReservedMoney    int64 `json:"reserved_money"`
	AvailableBalance int64 `json:"available_balance"`
	DailyLimit       int64 `json:"daily_limit"`
	DailySpent       int64 `json:"daily_spent"`
}

func (a *App) GetWallet(ctx context.Context, guildID int64) (GetWalletResponse, error) {
	g, err := a.repos.Guild.GetByID(ctx, guildID)
	if err != nil {
		return GetWalletResponse{}, fmt.Errorf("get wallet: %w", err)
	}

	return GetWalletResponse{
		GuildID:          g.ID,
		TotalMoney:       g.TotalMoney,
		ReservedMoney:    g.ReservedMoney,
		AvailableBalance: g.AvailableBalance(),
		DailyLimit:       g.DailyLimit,
		DailySpent:       g.DailySpent,
	}, nil
}
