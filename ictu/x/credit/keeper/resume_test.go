package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "ictu/testutil/keeper"
	"ictu/testutil/nullify"
	"ictu/x/credit/keeper"
	"ictu/x/credit/types"
)

func createTestResume(keeper *keeper.Keeper, ctx sdk.Context) types.Resume {
	item := types.Resume{}
	keeper.SetResume(ctx, item)
	return item
}

func TestResumeGet(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	item := createTestResume(keeper, ctx)
	rst, found := keeper.GetResume(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestResumeRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	createTestResume(keeper, ctx)
	keeper.RemoveResume(ctx)
	_, found := keeper.GetResume(ctx)
	require.False(t, found)
}
