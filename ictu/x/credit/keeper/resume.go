package keeper

import (
	"ictu/x/credit/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateResumeNewBalance(ctx sdk.Context, balance types.Balance) {
	resume, found := k.GetResume(ctx)
	if !found {
		panic(types.PanicResumeCantBeFound)
	}
	resume.CreditedTotal += balance.Credited
	resume.CreditedCurrent += balance.Credited
	resume.UpdateBalances()
	k.SetResume(ctx, resume)
}

// incorrect implementation with the returns
func (k Keeper) UpdateResumeUpdateBalance(ctx sdk.Context, balance types.Balance) {
	resume, found := k.GetResume(ctx)
	if !found {
		panic(types.PanicResumeCantBeFound)
	}
	resume.ReturnedTotal += balance.Returned
	resume.ReturnedCurrent += balance.Returned
	resume.UpdateBalances()
	k.SetResume(ctx, resume)
}

func (k Keeper) UpdateResumeRemoveBalance(ctx sdk.Context, balance types.Balance) {
	resume, found := k.GetResume(ctx)
	if !found {
		panic(types.PanicResumeCantBeFound)
	}
	resume.CreditedCurrent -= balance.Credited
	resume.ReturnedCurrent -= balance.Returned
	resume.UpdateBalances()
	k.SetResume(ctx, resume)
}

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
