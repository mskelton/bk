package main

import (
	"os"

	"github.com/mskelton/bk/cmd"
)

func main() {
	cmd := cmd.NewCmdRoot()
	err := cmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
