package tictactoe_test

import (
	"testing"

	keepertest "github.com/shifty11/tic-tac-toe/testutil/keeper"
	"github.com/shifty11/tic-tac-toe/testutil/nullify"
	"github.com/shifty11/tic-tac-toe/x/tictactoe"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextGameId: 7,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, *k, genesisState)
	got := tictactoe.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
