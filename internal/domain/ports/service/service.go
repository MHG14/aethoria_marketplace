package service

import (
	"context"
	"time"
)

type PriceOracle interface {
	GetBasePrice(ctx context.Context, itemID int64) (int64, error)
}

type Clock interface {
	Now() time.Time
}

type Services struct {
	Oracle PriceOracle
	Clock  Clock
}
