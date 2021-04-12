package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo to youer list",
	Long:  `Create a new todo to your list to track`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
