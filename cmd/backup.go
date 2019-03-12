package cmd

import (
	"fmt"
	"github.com/nwtgck/ibk"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
)

var useLocalTime bool
var backupDirPath string

func init() {
	RootCmd.AddCommand(backupCmd)
	backupCmd.Flags().BoolVar(&useLocalTime, "local-time", false, "use local time not UTC")
	backupCmd.Flags().StringVarP(&backupDirPath, "backup-path", "b", "", "backup destination path (default: <src>.ibk)")
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			// NOTE: This interface of CLI can be change
			return fmt.Errorf("Source path is required\n")
		}

		// Get source directory path
		srcDirPath := args[0]

		// Get current local time
		now := time.Now()
		// It not using local-time
		if !useLocalTime {
			// Set UTC because default time is UTC
			now = now.UTC()
		}

		// If backupDirPath is not set
		if backupDirPath == "" {
			// Get the base name of the source directory
			srcBaseName := filepath.Base(srcDirPath)
			// Assign default backup directory name
			backupDirPath = fmt.Sprintf("%s.ibk", srcBaseName)
		}

		// Run backup
		err := ibk.Backup(srcDirPath, backupDirPath, now)
		return err
	},
}
