package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/item"
)

type CreateItemRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	OwnerID int64  `json:"owner_id"`
}

type CreateItemResponse struct {
	Item item.Item `json:"item"`
}

func (a *App) CreateItem(ctx context.Context, req CreateItemRequest) (CreateItemResponse, error) {
	if req.Name == "" {
		return CreateItemResponse{}, fmt.Errorf("%w: name is required", ErrInvalidInput)
	}

	itemType := item.Type(req.Type)
	switch itemType {
	case item.Common, item.Rare, item.Legendary:
	default:
		return CreateItemResponse{}, fmt.Errorf("%w: invalid item type", ErrInvalidInput)
	}

	// verify owner exists
	if _, err := a.repos.Guild.GetByID(ctx, req.OwnerID); err != nil {
		return CreateItemResponse{}, fmt.Errorf("owner guild not found: %w", ErrNotFound)
	}

	i, err := a.repos.Item.Create(ctx, item.Item{
		Name:    req.Name,
		Type:    itemType,
		Status:  item.Available,
		OwnerID: req.OwnerID,
	})
	if err != nil {
		return CreateItemResponse{}, fmt.Errorf("create item: %w", err)
	}

	return CreateItemResponse{Item: i}, nil
}
