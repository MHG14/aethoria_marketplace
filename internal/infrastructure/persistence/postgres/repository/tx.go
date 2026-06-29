package repository

import (
	"context"
	"fmt"

	domainrepo "github.com/MHG14/aethoria_marketplace/internal/domain/ports/repository"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txManager struct {
	pool *pgxpool.Pool
}

func (m *txManager) WithTx(ctx context.Context, fn func(ctx context.Context, repos domainrepo.Repositories) error) error {
	tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	q := db.New(tx)
	repos := domainrepo.Repositories{
		Item:    &itemRepo{q: q},
		Guild:   &guildRepo{q: q},
		Listing: &listingRepo{q: q},
		Auction: &auctionRepo{q: q},
		Bid:     &bidRepo{q: q},
		Trade:   &tradeRepo{q: q},
		Wallet:  &walletRepo{q: q},
	}

	if err := fn(ctx, repos); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
