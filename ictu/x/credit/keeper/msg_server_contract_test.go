package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "ictu/testutil/keeper"
	"ictu/x/credit/keeper"
	"ictu/x/credit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestContractMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CreditKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateContract{Creator: creator,
			Req:  strconv.Itoa(i),
			Prov: strconv.Itoa(i),
		}
		_, err := srv.CreateContract(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetContract(ctx,
			"",
			expected.Req,
			expected.Prov,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestContractMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateContract
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateContract{Creator: creator,
				Uid:  strconv.Itoa(0),
				Req:  strconv.Itoa(0),
				Prov: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateContract{Creator: "B",
				Uid:  strconv.Itoa(0),
				Req:  strconv.Itoa(0),
				Prov: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateContract{Creator: creator,
				Uid:  strconv.Itoa(100000),
				Req:  strconv.Itoa(100000),
				Prov: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CreditKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateContract{Creator: creator,
				Req:  strconv.Itoa(0),
				Prov: strconv.Itoa(0),
			}
			_, err := srv.CreateContract(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateContract(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetContract(ctx,
					"",
					expected.Req,
					expected.Prov,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestContractMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteContract
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteContract{Creator: creator,
				Uid:  strconv.Itoa(0),
				Req:  strconv.Itoa(0),
				Prov: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteContract{Creator: "B",
				Uid:  strconv.Itoa(0),
				Req:  strconv.Itoa(0),
				Prov: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteContract{Creator: creator,
				Uid:  strconv.Itoa(100000),
				Req:  strconv.Itoa(100000),
				Prov: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CreditKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateContract(wctx, &types.MsgCreateContract{Creator: creator,
				Req:  strconv.Itoa(0),
				Prov: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteContract(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetContract(ctx,
					tc.request.Uid,
					tc.request.Req,
					tc.request.Prov,
				)
				require.False(t, found)
			}
		})
	}
}
