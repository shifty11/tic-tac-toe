package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAcceptInvite(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	acceptResponse, err := msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: bob,
		GameId:  1,
	})
	require.Nil(t, err)
	require.EqualValues(t, &types.MsgAcceptInviteResponse{}, acceptResponse)

	game, found := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:   "1",
		Player1: bob,
		Player2: alice,
		Turn:    1,
		Status:  types.StoredGame_IN_PROGRESS,
		Winner:  types.StoredGame_NONE,
		Board:   ".........",
	}, game)

	ctx := sdk.UnwrapSDKContext(context)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[1]
	require.EqualValues(t, "shifty11.tictactoe.tictactoe.EventInviteAccepted", event.Type)
	require.Len(t, event.Attributes, 3)
	require.Equal(t, event.Attributes[0].String(), `gameIndex: "1"`)
	require.Equal(t, event.Attributes[1].String(), fmt.Sprintf(`player1: "%s"`, bob))
	require.Equal(t, event.Attributes[2].String(), fmt.Sprintf(`player2: "%s"`, alice))
}

func TestAcceptInviteNotInvited(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	acceptResponse, err := msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: carol,
		GameId:  1,
	})
	require.Nil(t, acceptResponse)
	require.Equal(t, err.Error(), types.ErrPlayerNotInvited.Error()+fmt.Sprintf(": player %s is not part of game with id 1", carol))
}

func TestAcceptInviteGameDoesNoteExist(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	acceptResponse, err := msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: alice,
		GameId:  2,
	})
	require.Nil(t, acceptResponse)
	require.Equal(t, err.Error(), types.ErrGameNotFound.Error()+": game with id: 2 not found")
}

func TestAcceptInviteGameAlreadyStarted(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)

	acceptResponse, err := msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: alice,
		GameId:  1,
	})
	require.NoError(t, err)
	require.NotNil(t, acceptResponse)

	acceptResponse, err = msgServer.AcceptInvite(context, &types.MsgAcceptInvite{
		Creator: alice,
		GameId:  1,
	})
	require.Nil(t, acceptResponse)
	require.Equal(t, err.Error(), types.ErrGameAlreadyStarted.Error()+": game with id: 1 is already started")
}
