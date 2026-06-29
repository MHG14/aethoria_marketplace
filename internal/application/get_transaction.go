package application

import (
	"context"
	"fmt"

	"github.com/MHG14/aethoria_marketplace/internal/domain/wallet"
)

type GetTransactionsResponse struct {
	Transactions []wallet.Transaction `json:"transactions"`
}

func (a *App) GetTransactions(ctx context.Context, guildID int64) (GetTransactionsResponse, error) {
	txs, err := a.repos.Wallet.ListTransactions(ctx, guildID)
	if err != nil {
		return GetTransactionsResponse{}, fmt.Errorf("get transactions: %w", err)
	}
	return GetTransactionsResponse{Transactions: txs}, nil
}
