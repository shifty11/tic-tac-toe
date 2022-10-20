package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func TestDefineStartPlayer(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	bobAddress, err2 := sdk.AccAddressFromBech32(bob)
	p1, p2 := defineStartPlayer(aliceAddress, bobAddress)
	require.Equal(t, aliceAddress, p1)
	require.Equal(t, bobAddress, p2)
	require.Nil(t, err2)
	require.Nil(t, err1)

	//TODO: Test for hash generation starting with "0" -> create a function to create some random addresses and print them
}
