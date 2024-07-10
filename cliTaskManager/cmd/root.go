// package cmd

// import "github.com/spf13/cobra"

// var RootCmd = &cobra.Command{
// 	Use:   "task",
// 	Short: "Task is a CLI task manager",
// }

package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cliTaskManager",
	Short: "cliTaskManager is a CLI task manager",
}
