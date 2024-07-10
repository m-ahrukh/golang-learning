package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks of your list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list Command Called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
