package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ictu/x/credit/types"
)

// SetResume set resume in the store
func (k Keeper) SetResume(ctx sdk.Context, resume types.Resume) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResumeKey))
	b := k.cdc.MustMarshal(&resume)
	store.Set([]byte{0}, b)
}

// GetResume returns resume
func (k Keeper) GetResume(ctx sdk.Context) (val types.Resume, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResumeKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveResume removes resume from the store
func (k Keeper) RemoveResume(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResumeKey))
	store.Delete([]byte{0})
}
