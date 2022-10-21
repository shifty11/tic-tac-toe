package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/shifty11/tic-tac-toe/x/tictactoe/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAcceptInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-invite [game-id]",
		Short: "Broadcast message acceptInvite",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGameId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptInvite(
				clientCtx.GetFromAddress().String(),
				argGameId,
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
