package keeper

import (
	"context"
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"strconv"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextGameId, 10)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(err, types.ErrPlayerAddressInvalid.Error(), msg.Creator)
	}

	otherPlayer, err := sdk.AccAddressFromBech32(msg.Player2)
	if err != nil {
		return nil, errors.Wrapf(err, types.ErrPlayerAddressInvalid.Error(), msg.Player2)
	}

	storedGame := types.NewStoredGame(creator, otherPlayer, newIndex)
	k.Keeper.SetStoredGame(ctx, storedGame)

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventCreateGame{
			GameIndex: newIndex,
			Player1:   storedGame.Player1,
			Player2:   storedGame.Player2,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateGameResponse{GameId: newIndex}, nil
}
