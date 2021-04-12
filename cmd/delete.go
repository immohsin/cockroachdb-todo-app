package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Long:  `Delete a todo as complete after its done`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
