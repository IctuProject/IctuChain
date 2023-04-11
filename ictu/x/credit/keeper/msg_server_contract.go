package keeper

import (
	"context"
	"time"

	"ictu/x/credit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

func (k msgServer) CreateContract(goCtx context.Context, msg *types.MsgCreateContract) (*types.MsgCreateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Generate uuid for new contract
	uuid := uuid.New().String()

	// Check if the value already exists
	_, isFound := k.GetContract(
		ctx,
		uuid,
		msg.Req,
		msg.Prov,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var contract = types.Contract{
		Creator:          msg.Creator,
		Uid:              uuid,
		Req:              msg.Req,
		Prov:             msg.Prov,
		Amount:           msg.Amount,
		Desc:             msg.Desc,
		UtilLife:         msg.UtilLife,
		ReqSignature:     msg.Req,
		ProvSignature:    msg.Prov,
		IsExtension:      false,
		TimeCreated:      time.Now().String(),
		TimeReqAccepted:  "",
		TimeProvAccepted: "",
	}

	k.SetContract(
		ctx,
		contract,
	)
	return &types.MsgCreateContractResponse{}, nil
}

func (k msgServer) UpdateContract(goCtx context.Context, msg *types.MsgUpdateContract) (*types.MsgUpdateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetContract(
		ctx,
		msg.Uid,
		msg.Req,
		msg.Prov,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var contract = types.Contract{
		Creator:          msg.Creator,
		Uid:              msg.Uid,
		Req:              msg.Req,
		Prov:             msg.Prov,
		Amount:           msg.Amount,
		Desc:             msg.Desc,
		UtilLife:         msg.UtilLife,
		ReqSignature:     msg.ReqSignature,
		ProvSignature:    msg.ProvSignature,
		IsExtension:      msg.IsExtension,
		TimeCreated:      msg.TimeCreated,
		TimeReqAccepted:  msg.TimeReqAccepted,
		TimeProvAccepted: msg.TimeProvAccepted,
	}

	k.SetContract(ctx, contract)

	return &types.MsgUpdateContractResponse{}, nil
}

func (k msgServer) DeleteContract(goCtx context.Context, msg *types.MsgDeleteContract) (*types.MsgDeleteContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetContract(
		ctx,
		msg.Uid,
		msg.Req,
		msg.Prov,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveContract(
		ctx,
		msg.Uid,
		msg.Req,
		msg.Prov,
	)

	return &types.MsgDeleteContractResponse{}, nil
}
