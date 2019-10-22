package main

import (
	"fmt"

	"github.com/LambdaIM/lambda-debugtool/chain"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagPubkeyPrefix = "pubkey"
	FalgBechPrefix   = "bech"
)

func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show cons/real addr pubkey of valconspub address",
		RunE:  runShowCmd,
	}

	cmd.Flags().String(FlagPubkeyPrefix, "", "The pubKey of lambda val account")
	cmd.Flags().String(FalgBechPrefix, "pub", "The Bech32 prefix encoding for a key (pub|cons)")

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}

	return cmd
}

func runShowCmd(cmd *cobra.Command, args []string) error {
	consPubkey := viper.GetString(FlagPubkeyPrefix)
	if consPubkey == "" {
		return cmd.Help()
	}

	pubkey, err := sdk.GetConsPubKeyBech32(consPubkey)
	if err != nil {
		return err
	}

	bech := viper.GetString(FalgBechPrefix)
	switch bech {
	case "pub":
		{
			fmt.Println(pubkey.Address())
		}
	case "cons":
		{
			addr := pubkey.Address().Bytes()
			consAddr := sdk.ConsAddress(addr)
			fmt.Println(consAddr.String())
		}
	default:
		return fmt.Errorf("unknown bech type %v", bech)
	}

	return nil
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lambda debugtool",
		Short: "lambda debugtool",
	}

	return cmd
}

func main() {
	chain.SealConfig()

	root := rootCmd()
	root.AddCommand(showCmd())

	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}
