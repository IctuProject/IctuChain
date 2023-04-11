package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		BalanceList: []Balance{},
		Resume: &Resume{
			CreditedTotal:   0,
			ReturnedTotal:   0,
			BalanceTotal:    0,
			ReturnedCurrent: 0,
			BalanceCurrent:  0,
			CreditedCurrent: 0,
		},
		ContractList: []Contract{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in balance
	balanceIndexMap := make(map[string]struct{})

	for _, elem := range gs.BalanceList {
		index := string(BalanceKey(elem.Uid, elem.IdContract, elem.Requester))
		if _, ok := balanceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for balance")
		}
		balanceIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in contract
	contractIndexMap := make(map[string]struct{})

	for _, elem := range gs.ContractList {
		index := string(ContractKey(elem.Uid, elem.Req, elem.Prov))
		if _, ok := contractIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for contract")
		}
		contractIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
