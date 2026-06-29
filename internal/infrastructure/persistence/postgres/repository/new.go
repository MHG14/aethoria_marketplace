package repository

import (
	domainrepo "github.com/MHG14/aethoria_marketplace/internal/domain/ports/repository"
	db "github.com/MHG14/aethoria_marketplace/internal/infrastructure/persistence/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type itemRepo struct{ q *db.Queries }
type guildRepo struct{ q *db.Queries }
type listingRepo struct{ q *db.Queries }
type auctionRepo struct{ q *db.Queries }
type bidRepo struct{ q *db.Queries }
type tradeRepo struct{ q *db.Queries }
type walletRepo struct{ q *db.Queries }

func New(pool *pgxpool.Pool) domainrepo.Repositories {
	q := db.New(pool)
	return domainrepo.Repositories{
		Item:      &itemRepo{q: q},
		Guild:     &guildRepo{q: q},
		Listing:   &listingRepo{q: q},
		Auction:   &auctionRepo{q: q},
		Bid:       &bidRepo{q: q},
		Trade:     &tradeRepo{q: q},
		Wallet:    &walletRepo{q: q},
		TxManager: &txManager{pool: pool},
	}
}
