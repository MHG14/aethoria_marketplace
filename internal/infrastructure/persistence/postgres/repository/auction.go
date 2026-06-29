package repository

import (
	"context"
	"time"

	"github.com/MHG14/aethoria_marketplace/internal/domain/auction"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *auctionRepo) Create(ctx context.Context, a auction.Auction) (auction.Auction, error) {
	row, err := r.q.CreateAuction(ctx, &db.CreateAuctionParams{
		ItemID:        a.ItemID,
		SellerID:      a.SellerID,
		StartingPrice: a.StartingPrice,
		EndTime: pgtype.Timestamptz{
			Time:  a.EndTime,
			Valid: true,
		},
	})
	if err != nil {
		return auction.Auction{}, err
	}
	return toAuction(row), nil
}

func (r *auctionRepo) GetByID(ctx context.Context, id int64) (auction.Auction, error) {
	row, err := r.q.GetAuction(ctx, id)
	if err != nil {
		return auction.Auction{}, err
	}
	return toAuction(row), nil
}

func (r *auctionRepo) GetByIDForUpdate(ctx context.Context, id int64) (auction.Auction, error) {
	row, err := r.q.GetAuctionForUpdate(ctx, id)
	if err != nil {
		return auction.Auction{}, err
	}
	return toAuction(row), nil
}

func (r *auctionRepo) GetActiveByItemID(ctx context.Context, itemID int64) (auction.Auction, error) {
	row, err := r.q.GetActiveAuctionByItem(ctx, itemID)
	if err != nil {
		return auction.Auction{}, err
	}
	return toAuction(row), nil
}

func (r *auctionRepo) UpdateBid(ctx context.Context, id, highestBid int64, highestBidderID *int64, endTime time.Time) (auction.Auction, error) {
	row, err := r.q.UpdateAuctionBid(ctx, &db.UpdateAuctionBidParams{
		ID:              id,
		HighestBid:      highestBid,
		HighestBidderID: highestBidderID,
		EndTime:         pgtype.Timestamptz{Time: endTime, Valid: true},
	})
	if err != nil {
		return auction.Auction{}, err
	}
	return toAuction(row), nil
}

func (r *auctionRepo) UpdateStatus(ctx context.Context, id int64, status auction.Status) (auction.Auction, error) {
	row, err := r.q.UpdateAuctionStatus(ctx, &db.UpdateAuctionStatusParams{
		ID:     id,
		Status: db.AuctionStatus(status),
	})
	if err != nil {
		return auction.Auction{}, err
	}
	return toAuction(row), nil
}

func (r *auctionRepo) ListActive(ctx context.Context) ([]auction.Auction, error) {
	rows, err := r.q.ListActiveAuctions(ctx)
	if err != nil {
		return nil, err
	}
	return toAuctions(rows), nil
}

func (r *auctionRepo) ListExpired(ctx context.Context) ([]auction.Auction, error) {
	rows, err := r.q.ListExpiredAuctions(ctx)
	if err != nil {
		return nil, err
	}
	return toAuctions(rows), nil
}

func toAuction(row *db.Auction) auction.Auction {
	return auction.Auction{
		ID:              row.ID,
		ItemID:          row.ItemID,
		SellerID:        row.SellerID,
		StartingPrice:   row.StartingPrice,
		HighestBid:      row.HighestBid,
		HighestBidderID: row.HighestBidderID,
		EndTime:         row.EndTime.Time,
		OriginalEndTime: row.OriginalEndTime.Time,
		Status:          auction.Status(row.Status),
		CreatedAt:       row.CreatedAt.Time,
	}
}

func toAuctions(rows []*db.Auction) []auction.Auction {
	auctions := make([]auction.Auction, len(rows))
	for i, row := range rows {
		auctions[i] = toAuction(row)
	}
	return auctions
}
