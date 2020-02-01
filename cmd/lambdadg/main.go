package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/LambdaIM/lambda-debugtool/chain"
)

func preRun(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().StringP(FlagLambdaCliHome, "", DefaultLambdaCliHome, "config lambdacli home directory")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		return nil
	}

	return cmd
}

func main() {
	chain.SealConfig()

	root := rootCmd()
	root.AddCommand(showCmd())
	root.AddCommand(repairKeysCmd())
	root.SilenceUsage = true

	_ = preRun(root).Execute()
}
