package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"path"
)

const (
	flagKeys = "keys"
	flagData = "data"
)

func repairKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repair",
		Short: "repair db [--keys | --data <dbname> ]",
		RunE:  runRepairCmd,
	}
	cmd.Flags().BoolP(flagKeys, "k", false, "repair keys")
	cmd.Flags().BoolP(flagData, "d", false, "repair data <dbname>")
	return cmd
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func runRepairCmd(cmd *cobra.Command, args []string) error {
	if viper.GetBool(flagKeys) {
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
	} else if viper.GetBool(flagData) {
		if len(args) == 0 {
			return cmd.Help()
		}

		dbPath := args[0]

		if !FileExists(dbPath) {
			return fmt.Errorf("db: %v was not found", dbPath)
		}

		_, err := leveldb.OpenFile(dbPath, nil)
		if err != nil {
			if _, err = leveldb.RecoverFile(dbPath, nil); err != nil {
				return fmt.Errorf(`recoverfile %s : %s`, dbPath, err.Error())
			}
			fmt.Printf("repair successful, db: %v\n", dbPath)
		} else {
			fmt.Printf("no repair required, db: %v\n", dbPath)
		}
	} else {
		return cmd.Help()
	}

	return nil
}
