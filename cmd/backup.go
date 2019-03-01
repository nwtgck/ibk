package cmd

import (
	"fmt"
	"github.com/nwtgck/ibk"
	"github.com/spf13/cobra"
	"time"
)

var useLocalTime bool

func init() {
	RootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolVar(&useLocalTime, "local-time", false, "use local time not UTC")
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			// NOTE: This interface of CLI can be change
			return fmt.Errorf("Argument should be two\n")
		}

		// Get current local time
		now := time.Now()
		// It not using local-time
		if !useLocalTime {
			// Set UTC because default time is UTC
			now = now.UTC()
		}
		err := ibk.Backup(args[0], args[1], now)
		return err
	},
}
