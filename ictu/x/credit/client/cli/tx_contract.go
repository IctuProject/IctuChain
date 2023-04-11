package cli

import (
	"ictu/x/credit/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateContract() *cobra.Command {
	cmd := &cobra.Command{
		//Use:   "create-contract [uid] [req] [prov] [amount] [desc] [util-life] [req-signature] [prov-signature] [is-extension] [time-created] [time-req-accepted] [time-prov-accepted]",
		Use:   "create-contract [uid] [req] [prov] [amount] [desc] [util-life] [req-signature] [prov-signature]",
		Short: "Create a new contract",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexUid := args[0]
			indexReq := args[1]
			indexProv := args[2]

			// Get value arguments
			argAmount, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argDesc := args[4]
			argUtilLife := args[5]
			argReqSignature := args[6]
			argProvSignature := args[7]
			argIsExtension, err := cast.ToBoolE(args[8])
			if err != nil {
				return err
			}
			argTimeCreated := args[9]
			argTimeReqAccepted := args[10]
			argTimeProvAccepted := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateContract(
				clientCtx.GetFromAddress().String(),
				indexUid,
				indexReq,
				indexProv,
				argAmount,
				argDesc,
				argUtilLife,
				argReqSignature,
				argProvSignature,
				argIsExtension,
				argTimeCreated,
				argTimeReqAccepted,
				argTimeProvAccepted,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-contract [uid] [req] [prov] [amount] [desc] [util-life] [req-signature] [prov-signature] [is-extension] [time-created] [time-req-accepted] [time-prov-accepted]",
		Short: "Update a contract",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexUid := args[0]
			indexReq := args[1]
			indexProv := args[2]

			// Get value arguments
			argAmount, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argDesc := args[4]
			argUtilLife := args[5]
			argReqSignature := args[6]
			argProvSignature := args[7]
			argIsExtension, err := cast.ToBoolE(args[8])
			if err != nil {
				return err
			}
			argTimeCreated := args[9]
			argTimeReqAccepted := args[10]
			argTimeProvAccepted := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateContract(
				clientCtx.GetFromAddress().String(),
				indexUid,
				indexReq,
				indexProv,
				argAmount,
				argDesc,
				argUtilLife,
				argReqSignature,
				argProvSignature,
				argIsExtension,
				argTimeCreated,
				argTimeReqAccepted,
				argTimeProvAccepted,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-contract [uid] [req] [prov]",
		Short: "Delete a contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexUid := args[0]
			indexReq := args[1]
			indexProv := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteContract(
				clientCtx.GetFromAddress().String(),
				indexUid,
				indexReq,
				indexProv,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
