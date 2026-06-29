package auction

import "errors"

var (
	ErrAuctionNotActive    = errors.New("auction is not active")
	ErrSellerCannotBid     = errors.New("seller cannot bid on own item")
	ErrBidTooLow           = errors.New("bid must be at least 5% above current highest")
	ErrCannotCancelTopBid  = errors.New("cannot cancel bid while you are the highest bidder")
	ErrActiveAuctionExists = errors.New("an active auction already exists for this item")
)
