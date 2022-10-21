package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/tictactoe module sentinel errors
var (
	ErrPlayerAddressInvalid = errors.Register(ModuleName, 1100, "player address is invalid")
	ErrGameNotFound         = errors.Register(ModuleName, 1101, "game not found")
	ErrPlayerNotPartOfGame  = errors.Register(ModuleName, 1102, "player is not part of the game")
	ErrGameAlreadyStarted   = errors.Register(ModuleName, 1103, "game already started")
	ErrGameNotInProgress    = errors.Register(ModuleName, 1104, "game not in progress")
	ErrNotPlayersTurn       = errors.Register(ModuleName, 1105, "not players turn")
	ErrInvalidMove          = errors.Register(ModuleName, 1106, "invalid move")
)
