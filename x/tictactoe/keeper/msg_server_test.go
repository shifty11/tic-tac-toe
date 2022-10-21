package keeper_test

import (
	"context"
	"github.com/shifty11/tic-tac-toe/x/tictactoe"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/shifty11/tic-tac-toe/testutil/keeper"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/keeper"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
