package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAcceptInvite = "accept_invite"

var _ sdk.Msg = &MsgAcceptInvite{}

func NewMsgAcceptInvite(creator string, gameId uint64) *MsgAcceptInvite {
	return &MsgAcceptInvite{
		Creator: creator,
		GameId:  gameId,
	}
}

func (msg *MsgAcceptInvite) Route() string {
	return RouterKey
}

func (msg *MsgAcceptInvite) Type() string {
	return TypeMsgAcceptInvite
}

func (msg *MsgAcceptInvite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptInvite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptInvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
