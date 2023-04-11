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

func createNContract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Contract {
	items := make([]types.Contract, n)
	for i := range items {
		items[i].Uid = strconv.Itoa(i)
		items[i].Req = strconv.Itoa(i)
		items[i].Prov = strconv.Itoa(i)

		keeper.SetContract(ctx, items[i])
	}
	return items
}

func TestContractGet(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNContract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetContract(ctx,
			item.Uid,
			item.Req,
			item.Prov,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestContractRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveContract(ctx,
			item.Uid,
			item.Req,
			item.Prov,
		)
		_, found := keeper.GetContract(ctx,
			item.Uid,
			item.Req,
			item.Prov,
		)
		require.False(t, found)
	}
}

func TestContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllContract(ctx)),
	)
}
