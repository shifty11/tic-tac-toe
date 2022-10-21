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

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
