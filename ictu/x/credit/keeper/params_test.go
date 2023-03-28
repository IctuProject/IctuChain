package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ictu/testutil/keeper"
	"ictu/x/credit/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CreditKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
