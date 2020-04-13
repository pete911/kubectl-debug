package flag

import "github.com/spf13/cobra"

type Pod struct {
	Namespace string
	Label     string
}

func InitPodFlags(cmd *cobra.Command, flags *Pod) {

	cmd.Flags().StringVarP(
		&flags.Namespace,
		"namespace",
		"n",
		"default",
		"kubernetes namespace",
	)
	cmd.Flags().StringVarP(
		&flags.Label,
		"label",
		"l",
		"",
		"kubernetes label",
	)
}
