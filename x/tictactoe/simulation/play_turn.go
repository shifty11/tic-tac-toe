package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/keeper"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
)

func SimulateMsgPlayTurn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPlayTurn{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the PlayTurn simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "PlayTurn simulation not implemented"), nil, nil
	}
}
