package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"ictu/x/credit/types"
)

func CmdListBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-balance",
		Short: "list all balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBalanceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BalanceAll(context.Background(), params)
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

func CmdShowBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-balance [uid] [id-contract] [requester]",
		Short: "shows a balance",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUid := args[0]
			argIdContract := args[1]
			argRequester := args[2]

			params := &types.QueryGetBalanceRequest{
				Uid:        argUid,
				IdContract: argIdContract,
				Requester:  argRequester,
			}

			res, err := queryClient.Balance(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
