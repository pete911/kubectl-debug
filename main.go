package main

import (
	"github.com/pete911/kubectl-debug/cmd"
	"os"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
