package cmd

import (
	"fmt"
	"github.com/pete911/kubectl-debug/cmd/flag"
	"github.com/pete911/kubectl-debug/pkg/debug"
	"github.com/pete911/kubectl-debug/util"
	"github.com/spf13/cobra"
)

var (
	cmdPod = &cobra.Command{
		Use:     "pod",
		Aliases: []string{"pods", "po"},
		Short:   "debug pod",
		Long:    "",
		RunE:    runPodCmd,
	}
)

func init() {
	RootCmd.AddCommand(cmdPod)
}

func runPodCmd(_ *cobra.Command, _ []string) error {

	restConfig, err := util.RestConfig(flag.KubeconfigPath)
	if err != nil {
		return fmt.Errorf("loading kubernetes restconfig: %w", err)
	}
	return debug.Pod(restConfig)
}
