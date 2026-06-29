package repository

import (
	"context"

	"github.com/MHG14/aethoria_marketplace/internal/domain/auction/bid"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
)

func (r *bidRepo) Create(ctx context.Context, b bid.Bid) (bid.Bid, error) {
	row, err := r.q.CreateBid(ctx, &db.CreateBidParams{
		AuctionID: b.AuctionID,
		GuildID:   b.GuildID,
		Amount:    b.Amount,
	})
	if err != nil {
		return bid.Bid{}, err
	}
	return toBid(row), nil
}

func (r *bidRepo) GetByID(ctx context.Context, id int64) (bid.Bid, error) {
	row, err := r.q.GetBid(ctx, id)
	if err != nil {
		return bid.Bid{}, err
	}
	return toBid(row), nil
}

func (r *bidRepo) Cancel(ctx context.Context, id int64) (bid.Bid, error) {
	row, err := r.q.CancelBid(ctx, id)
	if err != nil {
		return bid.Bid{}, err
	}
	return toBid(row), nil
}

func (r *bidRepo) ListByAuction(ctx context.Context, auctionID int64) ([]bid.Bid, error) {
	rows, err := r.q.ListBidsByAuction(ctx, auctionID)
	if err != nil {
		return nil, err
	}
	return toBids(rows), nil
}

func (r *bidRepo) ListActiveByGuildAndAuction(ctx context.Context, auctionID, guildID int64) ([]bid.Bid, error) {
	rows, err := r.q.ListActiveBidsByGuildAndAuction(ctx, &db.ListActiveBidsByGuildAndAuctionParams{
		AuctionID: auctionID,
		GuildID:   guildID,
	})
	if err != nil {
		return nil, err
	}
	return toBids(rows), nil
}

func toBid(row *db.Bid) bid.Bid {
	return bid.Bid{
		ID:          row.ID,
		AuctionID:   row.AuctionID,
		GuildID:     row.GuildID,
		Amount:      row.Amount,
		IsCancelled: row.IsCancelled,
		CreatedAt:   row.CreatedAt.Time,
	}
}

func toBids(rows []*db.Bid) []bid.Bid {
	bids := make([]bid.Bid, len(rows))
	for i, row := range rows {
		bids[i] = toBid(row)
	}
	return bids
}
