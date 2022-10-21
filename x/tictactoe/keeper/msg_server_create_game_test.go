package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameId: 1,
	}, *createResponse)

	game, found := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:   "1",
		Player1: bob,
		Player2: alice,
		Turn:    1,
		Status:  types.StoredGame_OPEN,
		Winner:  types.StoredGame_NONE,
		Board:   ".........",
	}, game)
	systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{NextGameId: 2}, systemInfo)

	ctx := sdk.UnwrapSDKContext(context)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, "shifty11.tictactoe.tictactoe.EventCreateGame", event.Type)
	require.Len(t, event.Attributes, 3)
	require.Equal(t, event.Attributes[0].String(), `gameIndex: "1"`)
	require.Equal(t, event.Attributes[1].String(), fmt.Sprintf(`player1: "%s"`, bob))
	require.Equal(t, event.Attributes[2].String(), fmt.Sprintf(`player2: "%s"`, alice))
}

func TestCreateGameWrongPlayerAddress(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: "wrong address",
	})
	require.Error(t, err)
	require.Nil(t, createResponse)
}

func TestCreate2Games(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameId: 1,
	}, *createResponse)

	expectedGame := types.StoredGame{
		Index:   "1",
		Player1: bob,
		Player2: alice,
		Turn:    1,
		Status:  types.StoredGame_OPEN,
		Winner:  types.StoredGame_NONE,
		Board:   ".........",
	}
	game, found := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found)
	require.EqualValues(t, expectedGame, game)
	systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{NextGameId: 2}, systemInfo)

	createResponse, err = msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameId: 2,
	}, *createResponse)

	game, found = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found)
	require.EqualValues(t, expectedGame, game)
	systemInfo, found = keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{NextGameId: 3}, systemInfo)
}
