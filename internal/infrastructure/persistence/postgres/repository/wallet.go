package repository

import (
	"context"

	"github.com/MHG14/aethoria_marketplace/internal/domain/wallet"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
)

func (r *walletRepo) CreateTransaction(ctx context.Context, tx wallet.Transaction) (wallet.Transaction, error) {
	row, err := r.q.CreateWalletTransaction(ctx, &db.CreateWalletTransactionParams{
		GuildID: tx.GuildID,
		Type:    db.TxType(tx.Type),
		Amount:  tx.Amount,
		RefType: db.RefType(tx.RefType),
		RefID:   int64(tx.RefID),
	})
	if err != nil {
		return wallet.Transaction{}, err
	}
	return toTransaction(row), nil
}

func (r *walletRepo) ListTransactions(ctx context.Context, guildID int64) ([]wallet.Transaction, error) {
	rows, err := r.q.ListWalletTransactions(ctx, guildID)
	if err != nil {
		return nil, err
	}
	txs := make([]wallet.Transaction, len(rows))
	for i, row := range rows {
		txs[i] = toTransaction(row)
	}
	return txs, nil
}

func toTransaction(row *db.WalletTransaction) wallet.Transaction {
	return wallet.Transaction{
		ID:        row.ID,
		GuildID:   row.GuildID,
		Type:      wallet.TxType(row.Type),
		Amount:    row.Amount,
		RefType:   wallet.RefType(row.RefType),
		RefID:     wallet.RefID(row.RefID),
		CreatedAt: row.CreatedAt.Time,
	}
}
