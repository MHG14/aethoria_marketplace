package guild

import "errors"

var (
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrDailyLimitExceeded = errors.New("daily limit exceeded")
)
