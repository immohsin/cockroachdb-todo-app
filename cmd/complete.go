package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a todo as complete",
	Long:  `Mark a todo as complete after its done`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
