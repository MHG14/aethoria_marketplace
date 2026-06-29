package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/item"
	"github.com/MHG14/aethoria_marketplace/internal/domain/listing"
	"github.com/MHG14/aethoria_marketplace/internal/domain/ports/repository"
)

type CreateListingRequest struct {
	ItemID   int64 `json:"item_id"`
	SellerID int64 `json:"seller_id"`
	Price    int64 `json:"price"`
}

type CreateListingResponse struct {
	Listing listing.Listing `json:"listing"`
}

func (a *App) CreateListing(ctx context.Context, req CreateListingRequest) (CreateListingResponse, error) {
	if req.Price <= 0 {
		return CreateListingResponse{}, fmt.Errorf("%w: price must be positive", ErrInvalidInput)
	}

	i, err := a.repos.Item.GetByID(ctx, req.ItemID)
	if err != nil {
		return CreateListingResponse{}, fmt.Errorf("item not found: %w", ErrNotFound)
	}
	if i.OwnerID != req.SellerID {
		return CreateListingResponse{}, fmt.Errorf("%w: you do not own this item", ErrForbidden)
	}
	if !i.CanBeListed() {
		if i.IsLegendary() {
			return CreateListingResponse{}, ErrLegendaryCannotBeListed
		}
		return CreateListingResponse{}, ErrItemNotAvailable
	}

	var l listing.Listing
	err = a.repos.TxManager.WithTx(ctx, func(ctx context.Context, repos repository.Repositories) error {
		i, err := repos.Item.GetByIDForUpdate(ctx, req.ItemID)
		if err != nil {
			return fmt.Errorf("item not found: %w", ErrNotFound)
		}
		if i.OwnerID != req.SellerID {
			return fmt.Errorf("%w: you do not own this item", ErrForbidden)
		}
		if !i.CanBeListed() {
			if i.IsLegendary() {
				return ErrLegendaryCannotBeListed
			}
			return ErrItemNotAvailable
		}

		l, err = repos.Listing.Create(ctx, listing.Listing{
			ItemID:   req.ItemID,
			SellerID: req.SellerID,
			Price:    req.Price,
			Status:   listing.Active,
		})
		if err != nil {
			return fmt.Errorf("create listing: %w", err)
		}

		if _, err = repos.Item.UpdateStatus(ctx, req.ItemID, item.Listed); err != nil {
			return fmt.Errorf("update item status: %w", err)
		}

		return nil
	})
	if err != nil {
		return CreateListingResponse{}, err
	}

	return CreateListingResponse{Listing: l}, nil
}
