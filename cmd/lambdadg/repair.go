package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"
	"path"
)

func repairKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repair",
		Short: "repair keys db",
		RunE:  runRepairCmd,
	}
	return cmd
}

func runRepairCmd(cmd *cobra.Command, args []string) error {
	keysPath := path.Join(viper.GetString(FlagLambdaCliHome), "keys", "keys.db")

	_, err := leveldb.OpenFile(keysPath, nil)
	if err != nil {
		if _, err = leveldb.RecoverFile(keysPath, nil); err != nil {
			return fmt.Errorf(`recoverfile %s : %s`, keysPath, err.Error())
		}
		fmt.Printf("repair successful, keys db: %v\n", keysPath)
	} else {
		fmt.Printf("no repair required, keys db: %v\n", keysPath)
	}

	return nil
}
