package cmd

import (
	"fmt"
	"github.com/nwtgck/ibk"
	"github.com/spf13/cobra"
)

var restoredPath string = "ibk_restored"

func init() {
	RootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVarP(&restoredPath, "restored-path", "r", restoredPath, fmt.Sprintf("restored destination path (default: %s)", restoredPath))
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			// NOTE: This interface of CLI can be change
			return fmt.Errorf("Backup directory path is required\n")
		}
		backupDirPath := args[0]
		// Restore
		err := ibk.Restore(backupDirPath, restoredPath)
		if err != nil {
			return err
		}
		fmt.Printf("Restored successfully in '%s'!\n", restoredPath)
		return nil
	},
}
