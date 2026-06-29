package handlers

import "github.com/MHG14/aethoria_marketplace/internal/application"

// Handlers bundles all handler groups.
// Passed as a single dependency to the server.
type Handlers struct {
	Item    *ItemHandler
	Listing *ListingHandler
	Auction *AuctionHandler
	Guild   *GuildHandler
}

func New(app *application.App) *Handlers {
	return &Handlers{
		Item:    NewItemHandler(app),
		Listing: NewListingHandler(app),
		Auction: NewAuctionHandler(app),
		Guild:   NewGuildHandler(app),
	}
}
