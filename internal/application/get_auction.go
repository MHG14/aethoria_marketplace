package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/auction"
)

type GetAuctionResponse struct {
	Auction auction.Auction `json:"auction"`
}

func (a *App) GetAuction(ctx context.Context, id int64) (GetAuctionResponse, error) {
	auc, err := a.repos.Auction.GetByID(ctx, id)
	if err != nil {
		return GetAuctionResponse{}, fmt.Errorf("get auction: %w", err)
	}
	return GetAuctionResponse{Auction: auc}, nil
}
