package httpserver

import (
	"errors"

	"github.com/MHG14/aethoria_marketplace/internal/domain/auction"
	domainerr "github.com/MHG14/aethoria_marketplace/internal/domain/error"
	"github.com/MHG14/aethoria_marketplace/internal/domain/guild"
	"github.com/MHG14/aethoria_marketplace/internal/domain/item"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "internal server error"

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(errorResponse{
			Error: fiberErr.Message,
		})
	}

	switch {
	case errors.Is(err, domainerr.ErrNotFound):
		code = fiber.StatusNotFound
		message = err.Error()
	case errors.Is(err, domainerr.ErrInvalidInput):
		code = fiber.StatusBadRequest
		message = err.Error()
	case errors.Is(err, guild.ErrInsufficientFunds):
		code = fiber.StatusUnprocessableEntity
		message = err.Error()
	case errors.Is(err, guild.ErrDailyLimitExceeded):
		code = fiber.StatusUnprocessableEntity
		message = err.Error()
	case errors.Is(err, auction.ErrAuctionNotActive):
		code = fiber.StatusConflict
		message = err.Error()
	case errors.Is(err, auction.ErrSellerCannotBid):
		code = fiber.StatusForbidden
		message = err.Error()
	case errors.Is(err, auction.ErrBidTooLow):
		code = fiber.StatusUnprocessableEntity
		message = err.Error()
	case errors.Is(err, auction.ErrCannotCancelTopBid):
		code = fiber.StatusForbidden
		message = err.Error()
	case errors.Is(err, item.ErrItemNotAvailable):
		code = fiber.StatusConflict
		message = err.Error()
	case errors.Is(err, item.ErrLegendaryCannotBeListed):
		code = fiber.StatusUnprocessableEntity
		message = err.Error()
	case errors.Is(err, auction.ErrActiveAuctionExists):
		code = fiber.StatusConflict
		message = err.Error()
	}

	return c.Status(code).JSON(errorResponse{Error: message})
}

type errorResponse struct {
	Error string `json:"error"`
}
