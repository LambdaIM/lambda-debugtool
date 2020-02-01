package main

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show cons/real addr pubkey of valconspub address",
		RunE:  runShowCmd,
	}

	cmd.Flags().String(FlagPubkeyPrefix, "", "The pubKey of lambda val account")
	cmd.Flags().String(FlagBechPrefix, "pub", "The Bech32 prefix encoding for a key (pub|cons)")

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

	bech := viper.GetString(FlagBechPrefix)
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
