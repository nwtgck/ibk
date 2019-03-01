package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func init() {
	cobra.OnInitialize()
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "ibk",
	Long:  "Increment Backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}