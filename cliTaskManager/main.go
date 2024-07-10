// package main

// import (
// 	"goLangLearning/cliTaskManager/cmd"
// )

// func main() {
// 	cmd.RootCmd.Execute()
// }

package main

import (
	"fmt"
	"goLangLearning/cliTaskManager/cmd"
	"goLangLearning/cliTaskManager/db"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	must(err)
	cmd.RootCmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
