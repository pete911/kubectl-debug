package flag

import (
	"github.com/pete911/kubectl-debug/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	KubeconfigPath string
)

func InitGlobal(cmd *cobra.Command) {

	homeDir, _ := util.GetHomeDir()
	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	cmd.PersistentFlags().StringVar(
		&KubeconfigPath,
		"kubeconfig",
		util.GetStringEnvVar("KUBECONFIG", kubeconfig),
		"path to kubeconfig file",
	)
}
