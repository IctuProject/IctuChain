package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "ictu/testutil/keeper"
	"ictu/testutil/nullify"
	"ictu/x/credit/types"
)

func TestResumeQuery(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestResume(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetResumeRequest
		response *types.QueryGetResumeResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetResumeRequest{},
			response: &types.QueryGetResumeResponse{Resume: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Resume(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
