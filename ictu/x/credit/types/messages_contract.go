package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateContract = "create_contract"
	TypeMsgUpdateContract = "update_contract"
	TypeMsgDeleteContract = "delete_contract"
)

var _ sdk.Msg = &MsgCreateContract{}

func NewMsgCreateContract(
	creator string,
	req string,
	prov string,
	amount uint64,
	desc string,
	utilLife uint64,

) *MsgCreateContract {
	return &MsgCreateContract{
		Creator:  creator,
		Req:      req,
		Prov:     prov,
		Amount:   amount,
		Desc:     desc,
		UtilLife: utilLife,
	}
}

func (msg *MsgCreateContract) Route() string {
	return RouterKey
}

func (msg *MsgCreateContract) Type() string {
	return TypeMsgCreateContract
}

func (msg *MsgCreateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateContract{}

func NewMsgUpdateContract(
	creator string,
	uid string,
	req string,
	prov string,
	amount uint64,
	desc string,
	utilLife uint64,
	reqSignature string,
	provSignature string,
	isExtension bool,
	timeCreated string,
	timeReqAccepted string,
	timeProvAccepted string,

) *MsgUpdateContract {
	return &MsgUpdateContract{
		Creator:          creator,
		Uid:              uid,
		Req:              req,
		Prov:             prov,
		Amount:           amount,
		Desc:             desc,
		UtilLife:         utilLife,
		ReqSignature:     reqSignature,
		ProvSignature:    provSignature,
		IsExtension:      isExtension,
		TimeCreated:      timeCreated,
		TimeReqAccepted:  timeReqAccepted,
		TimeProvAccepted: timeProvAccepted,
	}
}

func (msg *MsgUpdateContract) Route() string {
	return RouterKey
}

func (msg *MsgUpdateContract) Type() string {
	return TypeMsgUpdateContract
}

func (msg *MsgUpdateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteContract{}

func NewMsgDeleteContract(
	creator string,
	uid string,
	req string,
	prov string,

) *MsgDeleteContract {
	return &MsgDeleteContract{
		Creator: creator,
		Uid:     uid,
		Req:     req,
		Prov:    prov,
	}
}
func (msg *MsgDeleteContract) Route() string {
	return RouterKey
}

func (msg *MsgDeleteContract) Type() string {
	return TypeMsgDeleteContract
}

func (msg *MsgDeleteContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
