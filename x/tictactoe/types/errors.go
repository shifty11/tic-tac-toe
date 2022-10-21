package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/tictactoe module sentinel errors
var (
	ErrPlayerAddressInvalid = errors.Register(ModuleName, 1100, "player address is invalid")
	ErrGameNotFound         = errors.Register(ModuleName, 1101, "game not found")
	ErrPlayerNotInvited     = errors.Register(ModuleName, 1102, "player is not invited")
	ErrGameAlreadyStarted   = errors.Register(ModuleName, 1103, "game already started")
)
