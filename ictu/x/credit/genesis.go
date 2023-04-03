package credit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ictu/x/credit/keeper"
	"ictu/x/credit/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the balance
	for _, elem := range genState.BalanceList {
		k.SetBalance(ctx, elem)
	}
	// Set if defined
	if genState.Resume != nil {
		k.SetResume(ctx, *genState.Resume)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.BalanceList = k.GetAllBalance(ctx)
	// Get all resume
	resume, found := k.GetResume(ctx)
	if found {
		genesis.Resume = &resume
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
