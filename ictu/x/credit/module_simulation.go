package credit

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"ictu/testutil/sample"
	creditsimulation "ictu/x/credit/simulation"
	"ictu/x/credit/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = creditsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateContract = "op_weight_msg_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContract int = 100

	opWeightMsgUpdateContract = "op_weight_msg_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateContract int = 100

	opWeightMsgDeleteContract = "op_weight_msg_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteContract int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	creditGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ContractList: []types.Contract{
			{
				Creator: sample.AccAddress(),
				Uid:     "0",
				Req:     "0",
				Prov:    "0",
			},
			{
				Creator: sample.AccAddress(),
				Uid:     "1",
				Req:     "1",
				Prov:    "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&creditGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContract, &weightMsgCreateContract, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContract = defaultWeightMsgCreateContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContract,
		creditsimulation.SimulateMsgCreateContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateContract, &weightMsgUpdateContract, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateContract = defaultWeightMsgUpdateContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateContract,
		creditsimulation.SimulateMsgUpdateContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteContract, &weightMsgDeleteContract, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteContract = defaultWeightMsgDeleteContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteContract,
		creditsimulation.SimulateMsgDeleteContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
