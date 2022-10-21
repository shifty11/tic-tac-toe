package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"
	"errors"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
)

func (k msgServer) AcceptInvite(goCtx context.Context, msg *types.MsgAcceptInvite) (*types.MsgAcceptInviteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, strconv.FormatUint(msg.GameId, 10))
	if !found {
		err := errors.New(fmt.Sprintf("game with id: %v not found", strconv.FormatUint(msg.GameId, 10)))
		return nil, sdkerrors.Wrapf(err, types.ErrGameNotFound.Error())
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPlayerAddressInvalid.Error(), msg.Creator)
	}

	if storedGame.Player1 != creator.String() && storedGame.Player2 != creator.String() {
		err = errors.New(fmt.Sprintf("player %v is not part of game with id %v", msg.Creator, msg.GameId))
		return nil, sdkerrors.Wrapf(err, types.ErrPlayerNotPartOfGame.Error())
	}

	if storedGame.Status != types.StoredGame_OPEN {
		err = errors.New(fmt.Sprintf("game with id: %v is already started", storedGame.Index))
		return nil, sdkerrors.Wrapf(err, types.ErrGameAlreadyStarted.Error())
	}

	storedGame.Status = types.StoredGame_IN_PROGRESS
	k.Keeper.SetStoredGame(ctx, storedGame)

	ctx.GasMeter().ConsumeGas(types.AcceptGameGas, "Accept game") //TODO: test this

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventInviteAccepted{
			GameIndex: msg.GameId,
			Player1:   storedGame.Player1,
			Player2:   storedGame.Player2,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgAcceptInviteResponse{}, nil
}
