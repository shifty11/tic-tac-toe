package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/tictactoe module sentinel errors
var (
	ErrPlayerAddressInvalid = errors.Register(ModuleName, 1100, "player address is invalid: %v")
)
