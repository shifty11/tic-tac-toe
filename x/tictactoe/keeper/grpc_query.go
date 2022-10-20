package keeper

import (
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
