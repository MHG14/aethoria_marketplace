package repository

import (
	"context"

	"github.com/MHG14/aethoria_marketplace/internal/domain/trade"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
)

func (r *tradeRepo) Create(ctx context.Context, t trade.Trade) (trade.Trade, error) {
	row, err := r.q.CreateTrade(ctx, &db.CreateTradeParams{
		ItemID:   t.ItemID,
		SellerID: t.SellerID,
		BuyerID:  t.BuyerID,
		Price:    t.Price,
		Type:     db.TradeType(t.Type),
	})
	if err != nil {
		return trade.Trade{}, err
	}
	return toTrade(row), nil
}

func (r *tradeRepo) ListByGuild(ctx context.Context, guildID int64) ([]trade.Trade, error) {
	rows, err := r.q.ListTradesByGuild(ctx, guildID)
	if err != nil {
		return nil, err
	}
	trades := make([]trade.Trade, len(rows))
	for i, row := range rows {
		trades[i] = toTrade(row)
	}
	return trades, nil
}

func toTrade(row *db.Trade) trade.Trade {
	return trade.Trade{
		ID:        row.ID,
		ItemID:    row.ItemID,
		SellerID:  row.SellerID,
		BuyerID:   row.BuyerID,
		Price:     row.Price,
		Type:      trade.Type(row.Type),
		CreatedAt: row.CreatedAt.Time,
	}
}
