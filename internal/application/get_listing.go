package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/listing"
)

type GetListingResponse struct {
	Listing listing.Listing `json:"listing"`
}

func (a *App) GetListing(ctx context.Context, id int64) (GetListingResponse, error) {
	l, err := a.repos.Listing.GetByID(ctx, id)
	if err != nil {
		return GetListingResponse{}, fmt.Errorf("get listing: %w", err)
	}
	return GetListingResponse{Listing: l}, nil
}
