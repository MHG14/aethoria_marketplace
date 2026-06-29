package handlers

import (
	"github.com/MHG14/aethoria_marketplace/internal/application"
	"github.com/gofiber/fiber/v2"
)

type AuctionHandler struct {
	app *application.App
}

func NewAuctionHandler(app *application.App) *AuctionHandler {
	return &AuctionHandler{app: app}
}

func (h *AuctionHandler) Create(c *fiber.Ctx) error {
	var req application.CreateAuctionRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.app.CreateAuction(c.Context(), req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *AuctionHandler) List(c *fiber.Ctx) error {
	result, err := h.app.ListActiveAuctions(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *AuctionHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.app.GetAuction(c.Context(), int64(id))
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *AuctionHandler) PlaceBid(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	var req application.PlaceBidRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	req.AuctionID = int64(id)
	result, err := h.app.PlaceBid(c.Context(), req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *AuctionHandler) CancelBid(c *fiber.Ctx) error {
	auctionID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	bidID, err := c.ParamsInt("bid_id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	var req application.CancelBidRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	req.AuctionID = int64(auctionID)
	req.BidID = int64(bidID)
	if err := h.app.CancelBid(c.Context(), req); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
