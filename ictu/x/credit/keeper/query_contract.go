package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ictu/x/credit/types"
)

func (k Keeper) ContractAll(goCtx context.Context, req *types.QueryAllContractRequest) (*types.QueryAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contracts []types.Contract
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	contractStore := prefix.NewStore(store, types.KeyPrefix(types.ContractKeyPrefix))

	pageRes, err := query.Paginate(contractStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.Contract
		if err := k.cdc.Unmarshal(value, &contract); err != nil {
			return err
		}

		contracts = append(contracts, contract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractResponse{Contract: contracts, Pagination: pageRes}, nil
}

func (k Keeper) Contract(goCtx context.Context, req *types.QueryGetContractRequest) (*types.QueryGetContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetContract(
		ctx,
		req.Uid,
		req.Req,
		req.Prov,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetContractResponse{Contract: val}, nil
}
