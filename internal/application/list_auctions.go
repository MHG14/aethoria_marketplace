package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/auction"
)

type ListActiveAuctionsResponse struct {
	Auctions []auction.Auction `json:"auctions"`
}

func (a *App) ListActiveAuctions(ctx context.Context) (ListActiveAuctionsResponse, error) {
	auctions, err := a.repos.Auction.ListActive(ctx)
	if err != nil {
		return ListActiveAuctionsResponse{}, fmt.Errorf("list auctions: %w", err)
	}
	return ListActiveAuctionsResponse{Auctions: auctions}, nil
}
