package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"ictu/x/credit/types"
)

func CmdListContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contract",
		Short: "list all contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllContractRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ContractAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-contract [uid] [req] [prov]",
		Short: "shows a contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUid := args[0]
			argReq := args[1]
			argProv := args[2]

			params := &types.QueryGetContractRequest{
				Uid:  argUid,
				Req:  argReq,
				Prov: argProv,
			}

			res, err := queryClient.Contract(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
