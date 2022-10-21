package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: bob,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameId: "1",
	}, *createResponse)

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
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Player2: "wrong address",
	})
	require.Error(t, err)
	require.Nil(t, createResponse)
}
