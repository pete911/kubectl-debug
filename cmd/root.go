package cmd

import (
	"github.com/pete911/kubectl-debug/cmd/flag"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "debug",
		Short: "debug kubernetes services",
		Long:  "",
	}
)

func init() {
	flag.InitGlobal(RootCmd)
}
