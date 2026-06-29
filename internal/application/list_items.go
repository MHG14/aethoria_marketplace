package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/item"
)

type ListItemsResponse struct {
	Items []item.Item `json:"items"`
}

func (a *App) ListItems(ctx context.Context) (ListItemsResponse, error) {
	items, err := a.repos.Item.List(ctx, 0)
	if err != nil {
		return ListItemsResponse{}, fmt.Errorf("list items: %w", err)
	}
	return ListItemsResponse{Items: items}, nil
}
