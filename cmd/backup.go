package cmd

import (
	"fmt"
	"github.com/nwtgck/ibk"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			// NOTE: This interface of CLI can be change
			return fmt.Errorf("Argument should be two\n")
		}
		err := ibk.Backup(args[0], args[1])
		return err
	},
}
