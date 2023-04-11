package credit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ictu/testutil/keeper"
	"ictu/testutil/nullify"
	"ictu/x/credit"
	"ictu/x/credit/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BalanceList: []types.Balance{
			{
				Uid:        "0",
				IdContract: "0",
				Requester:  "0",
			},
			{
				Uid:        "1",
				IdContract: "1",
				Requester:  "1",
			},
		},
		Resume: &types.Resume{
			CreditedTotal:   51,
			ReturnedTotal:   81,
			BalanceTotal:    6,
			CreditedCurrent: 49,
			ReturnedCurrent: 68,
			BalanceCurrent:  68,
		},
		ContractList: []types.Contract{
			{
				Uid:  "0",
				Req:  "0",
				Prov: "0",
			},
			{
				Uid:  "1",
				Req:  "1",
				Prov: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CreditKeeper(t)
	credit.InitGenesis(ctx, *k, genesisState)
	got := credit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BalanceList, got.BalanceList)
	require.Equal(t, genesisState.Resume, got.Resume)
	require.ElementsMatch(t, genesisState.ContractList, got.ContractList)
	// this line is used by starport scaffolding # genesis/test/assert
}
