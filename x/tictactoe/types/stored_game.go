package types

import (
	"crypto/sha1"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/rules"
)

func defineStartPlayer(player1 sdk.AccAddress, player2 sdk.AccAddress) (sdk.AccAddress, sdk.AccAddress) {
	data := []byte(player1.String() + player2.String())
	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)
	sh := fmt.Sprintf("%x", bs)
	if string(sh[0]) == "0" {
		return player1, player2
	}
	return player2, player1
}

func NewStoredGame(creator sdk.AccAddress, otherPlayer sdk.AccAddress, newIndex string) StoredGame {
	player1, player2 := defineStartPlayer(creator, otherPlayer)
	game := rules.NewGame()
	storedGame := StoredGame{
		Index:   newIndex,
		Player1: player1.String(),
		Player2: player2.String(),
		Turn:    uint64(game.Turn()),
		Status:  StoredGame_OPEN,
		Winner:  StoredGame_NONE,
		Board:   game.Board.String(),
	}
	return storedGame
}
