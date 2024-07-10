package cmd

import (
	"fmt"
	"goLangLearning/cliTaskManager/db"
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

		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Task added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
