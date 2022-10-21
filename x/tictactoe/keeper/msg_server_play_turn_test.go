package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayTurnGameDoesNoteExist(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	response, err := msgServer.PlayTurn(context, &types.MsgPlayTurn{
		Creator: alice,
		GameId:  2,
		X:       0,
		Y:       0,
	})
	require.Nil(t, response)
	require.Equal(t, types.ErrGameNotFound.Error()+": game with id: 2 not found", err.Error())
}

func TestPlayTurnPlayerNotPartOfGame(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	response, err := msgServer.PlayTurn(context, &types.MsgPlayTurn{
		Creator: carol,
		GameId:  1,
		X:       0,
		Y:       0,
	})
	require.Nil(t, response)
	require.Equal(t, types.ErrPlayerNotPartOfGame.Error()+fmt.Sprintf(": player %v is not part of game with id %v", carol, 1), err.Error())
}

func TestPlayTurnGameNotInProgress(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	response, err := msgServer.PlayTurn(context, &types.MsgPlayTurn{
		Creator: alice,
		GameId:  1,
		X:       0,
		Y:       0,
	})
	require.Nil(t, response)
	require.Equal(t, types.ErrGameNotInProgress.Error()+": game with id: 1 is not in progress", err.Error())
}

func TestPlayTurnNotPlayerTurn(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: bob,
		GameId:  1,
	})

	response, err := msgServer.PlayTurn(context, &types.MsgPlayTurn{
		Creator: alice,
		GameId:  1,
		X:       0,
		Y:       0,
	})
	require.Nil(t, response)
	require.Equal(t, types.ErrNotPlayersTurn.Error()+fmt.Sprintf(": player %v cannot play as it is not his turn", alice), err.Error())
}

func TestPlayTurnSuccess(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: bob,
		GameId:  1,
	})

	response, err := msgServer.PlayTurn(context, &types.MsgPlayTurn{
		Creator: bob,
		GameId:  1,
		X:       0,
		Y:       0,
	})
	require.Nil(t, err)
	require.EqualValues(t, &types.MsgPlayTurnResponse{
		Status: types.MsgPlayTurnResponse_IN_PROGRESS,
		Winner: types.MsgPlayTurnResponse_NONE,
		Board:  "*........",
	}, response)

	game, found := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:   "1",
		Player1: bob,
		Player2: alice,
		Turn:    2,
		Status:  types.StoredGame_IN_PROGRESS,
		Winner:  types.StoredGame_NONE,
		Board:   "*........",
	}, game)

	ctx := sdk.UnwrapSDKContext(context)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 3)
	event := events[2]
	require.EqualValues(t, "shifty11.tictactoe.tictactoe.EventTurnPlayed", event.Type)
	require.Len(t, event.Attributes, 6)
	require.Equal(t, event.Attributes[0].String(), `board: "*........"`)
	require.Equal(t, event.Attributes[1].String(), `column: 0`)
	require.Equal(t, event.Attributes[2].String(), `gameIndex: "1"`)
	require.Equal(t, event.Attributes[3].String(), fmt.Sprintf(`playedBy: "%s"`, bob))
	require.Equal(t, event.Attributes[4].String(), `row: 0`)
	require.Equal(t, event.Attributes[5].String(), `winner: "NONE"`)
}

//TODO: Test a couple of games with more turns and check the winner
