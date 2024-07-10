package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task in your list",
	Run: func(cmd *cobra.Command, args []string) {
		// for i, arg := range args {
		// 	fmt.Println(i+1, ":", arg)
		// }

		task := strings.Join(args, " ")
		fmt.Printf("Task added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
