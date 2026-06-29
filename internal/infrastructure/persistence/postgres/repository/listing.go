package repository

import (
	"context"

	"github.com/MHG14/aethoria_marketplace/internal/domain/listing"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
)

func (r *listingRepo) Create(ctx context.Context, l listing.Listing) (listing.Listing, error) {
	row, err := r.q.CreateListing(ctx, &db.CreateListingParams{
		ItemID:   l.ItemID,
		SellerID: l.SellerID,
		Price:    l.Price,
	})
	if err != nil {
		return listing.Listing{}, err
	}
	return toListing(row), nil
}

func (r *listingRepo) GetByID(ctx context.Context, id int64) (listing.Listing, error) {
	row, err := r.q.GetListing(ctx, id)
	if err != nil {
		return listing.Listing{}, err
	}
	return toListing(row), nil
}

func (r *listingRepo) GetByIDForUpdate(ctx context.Context, id int64) (listing.Listing, error) {
	row, err := r.q.GetListingForUpdate(ctx, id)
	if err != nil {
		return listing.Listing{}, err
	}
	return toListing(row), nil
}

func (r *listingRepo) UpdateStatus(ctx context.Context, id int64, status listing.Status, buyerID *int64) (listing.Listing, error) {
	row, err := r.q.UpdateListingStatus(ctx, &db.UpdateListingStatusParams{
		ID:      id,
		Status:  db.ListingStatus(status),
		BuyerID: buyerID,
	})
	if err != nil {
		return listing.Listing{}, err
	}
	return toListing(row), nil
}

func (r *listingRepo) ListActive(ctx context.Context) ([]listing.Listing, error) {
	rows, err := r.q.ListActiveListings(ctx)
	if err != nil {
		return nil, err
	}
	listings := make([]listing.Listing, len(rows))
	for i, row := range rows {
		listings[i] = toListing(row)
	}
	return listings, nil
}

func toListing(row *db.Listing) listing.Listing {
	return listing.Listing{
		ID:        row.ID,
		ItemID:    row.ItemID,
		SellerID:  row.SellerID,
		BuyerID:   row.BuyerID,
		Price:     row.Price,
		Status:    listing.Status(row.Status),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
