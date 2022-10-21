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

func (m *StoredGame) IsPlayersTurn(player sdk.AccAddress) bool {
	return (m.Player1 == player.String() && m.Turn == 1) || (m.Player2 == player.String() && m.Turn == 2)
}

func (m *StoredGame) PlayTurn(x uint64, y uint64) error {
	game := rules.ParseGame(m.Board)
	err := game.PlayTurn(int(m.Turn), int(x), int(y))
	if err != nil {
		return err
	}
	m.Turn = uint64(game.Turn())
	if game.GameStatus == rules.FirstPlayerWin {
		m.Winner = StoredGame_PLAYER1
	} else if game.GameStatus == rules.SecondPlayerWin {
		m.Winner = StoredGame_PLAYER2
	} else if game.GameStatus == rules.GameDraw {
		m.Winner = StoredGame_DRAW
	}
	if m.Winner == StoredGame_NONE {
		m.Status = StoredGame_IN_PROGRESS
	} else {
		m.Status = StoredGame_FINISHED
	}
	m.Board = game.Board.String()
	return nil
}
