package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/item"
	"github.com/MHG14/aethoria_marketplace/internal/domain/listing"
	"github.com/MHG14/aethoria_marketplace/internal/domain/ports/repository"
)

type CancelListingRequest struct {
	ListingID int64 `json:"listing_id"`
	SellerID  int64 `json:"seller_id"`
}

func (a *App) CancelListing(ctx context.Context, req CancelListingRequest) error {
	l, err := a.repos.Listing.GetByID(ctx, req.ListingID)
	if err != nil {
		return fmt.Errorf("listing not found: %w", ErrNotFound)
	}
	if l.SellerID != req.SellerID {
		return fmt.Errorf("%w: you do not own this listing", ErrForbidden)
	}
	if l.Status != listing.Active {
		return fmt.Errorf("%w: listing is not active", ErrItemNotAvailable)
	}

	return a.repos.TxManager.WithTx(ctx, func(ctx context.Context, repos repository.Repositories) error {
		l, err := repos.Listing.GetByIDForUpdate(ctx, req.ListingID)
		if err != nil {
			return fmt.Errorf("listing not found: %w", ErrNotFound)
		}
		if l.Status != listing.Active {
			return fmt.Errorf("%w: listing is not active", ErrItemNotAvailable)
		}

		if _, err = repos.Listing.UpdateStatus(ctx, l.ID, listing.Cancelled, nil); err != nil {
			return fmt.Errorf("cancel listing: %w", err)
		}

		if _, err = repos.Item.UpdateStatus(ctx, l.ItemID, item.Available); err != nil {
			return fmt.Errorf("restore item status: %w", err)
		}

		return nil
	})
}
