package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"ictu/x/credit/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
					CreditedTotal:   33,
					ReturnedTotal:   97,
					BalanceTotal:    89,
					CreditedCurrent: 3,
					ReturnedCurrent: 46,
					BalanceCurrent:  61,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated balance",
			genState: &types.GenesisState{
				BalanceList: []types.Balance{
					{
						Uid:        "0",
						IdContract: "0",
						Requester:  "0",
					},
					{
						Uid:        "0",
						IdContract: "0",
						Requester:  "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
