package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/item"
)

type GetItemResponse struct {
	Item item.Item `json:"item"`
}

func (a *App) GetItem(ctx context.Context, id int64) (GetItemResponse, error) {
	i, err := a.repos.Item.GetByID(ctx, id)
	if err != nil {
		return GetItemResponse{}, fmt.Errorf("get item: %w", err)
	}
	return GetItemResponse{Item: i}, nil
}
