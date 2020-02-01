package main

import (
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lambda debugtool",
		Short: "lambda debugtool",
	}

	return cmd
}
