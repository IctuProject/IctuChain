package keeper

import (
	"ictu/x/credit/types"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

// SetContract set a specific contract in the store from its index
func (k Keeper) SetContract(ctx sdk.Context, contract types.Contract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKeyPrefix))
	b := k.cdc.MustMarshal(&contract)
	store.Set(types.ContractKey(
		contract.Uid,
		contract.Req,
		contract.Prov,
	), b)

	nowPlusUtilLife := time.Now().Add(time.Duration(contract.UtilLife) * time.Hour)
	k.NewBalance(ctx, types.Balance{
		Uid:               uuid.NewString(),
		IdContract:        contract.Uid,
		Requester:         contract.Req,
		Credited:          contract.Amount,
		Returned:          0,
		ReturnPeriodicity: "h",
		ExpirationDate:    nowPlusUtilLife.String(),
	})
}

// GetContract returns a contract from its index
func (k Keeper) GetContract(
	ctx sdk.Context,
	uid string,
	req string,
	prov string,

) (val types.Contract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKeyPrefix))

	b := store.Get(types.ContractKey(
		uid,
		req,
		prov,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContract removes a contract from the store
func (k Keeper) RemoveContract(
	ctx sdk.Context,
	uid string,
	req string,
	prov string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKeyPrefix))
	store.Delete(types.ContractKey(
		uid,
		req,
		prov,
	))
}

// GetAllContract returns all contract
func (k Keeper) GetAllContract(ctx sdk.Context) (list []types.Contract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Contract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
