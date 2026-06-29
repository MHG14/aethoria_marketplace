package handlers

import (
	"github.com/MHG14/aethoria_marketplace/internal/application"
	"github.com/gofiber/fiber/v2"
)

type ListingHandler struct {
	app *application.App
}

func NewListingHandler(app *application.App) *ListingHandler {
	return &ListingHandler{app: app}
}

func (h *ListingHandler) Create(c *fiber.Ctx) error {
	var req application.CreateListingRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.app.CreateListing(c.Context(), req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *ListingHandler) List(c *fiber.Ctx) error {
	result, err := h.app.ListActiveListings(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *ListingHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.app.GetListing(c.Context(), int64(id))
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *ListingHandler) Buy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	var req application.BuyItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	req.ListingID = int64(id)
	result, err := h.app.BuyItem(c.Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *ListingHandler) Cancel(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	var req application.CancelListingRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	req.ListingID = int64(id)
	if err := h.app.CancelListing(c.Context(), req); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
