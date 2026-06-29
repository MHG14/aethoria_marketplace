package item

import "errors"

var (
	ErrItemNotAvailable        = errors.New("item is not available")
	ErrLegendaryCannotBeListed = errors.New("legendary items cannot be listed, use auction instead")
)
