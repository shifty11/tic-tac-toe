package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlayTurn = "play_turn"

var _ sdk.Msg = &MsgPlayTurn{}

func NewMsgPlayTurn(creator string, gameId uint64, x uint64, y uint64) *MsgPlayTurn {
	return &MsgPlayTurn{
		Creator: creator,
		GameId:  gameId,
		X:       x,
		Y:       y,
	}
}

func (msg *MsgPlayTurn) Route() string {
	return RouterKey
}

func (msg *MsgPlayTurn) Type() string {
	return TypeMsgPlayTurn
}

func (msg *MsgPlayTurn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlayTurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPlayTurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
