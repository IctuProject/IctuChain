package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ictu/testutil/keeper"
	"ictu/testutil/nullify"
	"ictu/x/credit/keeper"
	"ictu/x/credit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBalance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Balance {
	items := make([]types.Balance, n)
	for i := range items {
		items[i].Uid = strconv.Itoa(i)
		items[i].IdContract = strconv.Itoa(i)
		items[i].Requester = strconv.Itoa(i)

		keeper.SetBalance(ctx, items[i])
	}
	return items
}

func TestBalanceGet(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNBalance(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBalance(ctx,
			item.Uid,
			item.IdContract,
			item.Requester,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBalanceRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNBalance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBalance(ctx,
			item.Uid,
			item.IdContract,
			item.Requester,
		)
		_, found := keeper.GetBalance(ctx,
			item.Uid,
			item.IdContract,
			item.Requester,
		)
		require.False(t, found)
	}
}

func TestBalanceGetAll(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNBalance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBalance(ctx)),
	)
}
