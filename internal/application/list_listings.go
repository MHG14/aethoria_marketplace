package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/listing"
)

type ListActiveListingsResponse struct {
	Listings []listing.Listing `json:"listings"`
}

func (a *App) ListActiveListings(ctx context.Context) (ListActiveListingsResponse, error) {
	listings, err := a.repos.Listing.ListActive(ctx)
	if err != nil {
		return ListActiveListingsResponse{}, fmt.Errorf("list listings: %w", err)
	}
	return ListActiveListingsResponse{Listings: listings}, nil
}
