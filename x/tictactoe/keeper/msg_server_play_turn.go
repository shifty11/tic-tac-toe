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

func (k msgServer) PlayTurn(goCtx context.Context, msg *types.MsgPlayTurn) (*types.MsgPlayTurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, strconv.FormatUint(msg.GameId, 10))
	if !found {
		err := errors.New(fmt.Sprintf("game with id: %v not found", strconv.FormatUint(msg.GameId, 10)))
		return nil, sdkerrors.Wrapf(err, types.ErrGameNotFound.Error())
	}

	player, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrPlayerAddressInvalid.Error(), msg.Creator)
	}

	if storedGame.Player1 != player.String() && storedGame.Player2 != player.String() {
		err = errors.New(fmt.Sprintf("player %v is not part of game with id %v", msg.Creator, msg.GameId))
		return nil, sdkerrors.Wrapf(err, types.ErrPlayerNotPartOfGame.Error())
	}

	if storedGame.Status != types.StoredGame_IN_PROGRESS {
		err = errors.New(fmt.Sprintf("game with id: %v is not in progress", storedGame.Index))
		return nil, sdkerrors.Wrapf(err, types.ErrGameNotInProgress.Error())
	}

	if !storedGame.IsPlayersTurn(player) {
		err = errors.New(fmt.Sprintf("player %v cannot play as it is not his turn", msg.Creator))
		return nil, sdkerrors.Wrapf(err, types.ErrNotPlayersTurn.Error())
	}

	err = storedGame.PlayTurn(msg.X, msg.Y)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrInvalidMove.Error())
	}

	k.Keeper.SetStoredGame(ctx, storedGame)

	ctx.GasMeter().ConsumeGas(types.PlayTurnGas, "Play turn") //TODO: test this

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventTurnPlayed{
			GameIndex: msg.GameId,
			PlayedBy:  player.String(),
			Row:       uint32(msg.X),
			Column:    uint32(msg.Y),
			Winner:    types.EventTurnPlayed_WinnerStatus(storedGame.Winner),
			Board:     storedGame.Board,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgPlayTurnResponse{
		Status: types.MsgPlayTurnResponse_GameStatus(storedGame.Status),
		Winner: types.MsgPlayTurnResponse_WinnerStatus(storedGame.Winner),
		Board:  storedGame.Board,
	}, nil
}
