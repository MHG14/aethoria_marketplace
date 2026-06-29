package application

import "errors"

var (
	ErrNotFound                = errors.New("not found")
	ErrInvalidInput            = errors.New("invalid input")
	ErrForbidden               = errors.New("forbidden")
	ErrInsufficientFunds       = errors.New("insufficient funds")
	ErrDailyLimitExceeded      = errors.New("daily limit exceeded")
	ErrAuctionNotActive        = errors.New("auction is not active")
	ErrSellerCannotBid         = errors.New("seller cannot bid on own item")
	ErrBidTooLow               = errors.New("bid must be at least 5% above current highest")
	ErrCannotCancelTopBid      = errors.New("cannot cancel bid while you are the highest bidder")
	ErrItemNotAvailable        = errors.New("item is not available")
	ErrLegendaryCannotBeListed = errors.New("legendary items cannot be listed, use auction instead")
	ErrActiveAuctionExists     = errors.New("an active auction already exists for this item")
)
