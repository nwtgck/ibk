package cmd

import (
	"fmt"
	"github.com/nwtgck/ibk"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			// NOTE: This interface of CLI can be change
			return fmt.Errorf("Argument should be one\n")
		}
		err := ibk.Restore(args[0])
		return err
	},
}
