package cmd

import (
	"context"
	"fmt"
	"github.com/pete911/kubectl-debug/cmd/flag"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var (
	cmdPod = &cobra.Command{
		Use:     "pod",
		Aliases: []string{"pods", "po"},
		Short:   "debug pod",
		Long:    "",
		RunE:    runPodCmd,
	}
	podFlags flag.Pod
	stdOut   = os.Stdout
)

func init() {

	RootCmd.AddCommand(cmdPod)
	flag.InitPodFlags(cmdPod, &podFlags)
}

func runPodCmd(_ *cobra.Command, args []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	d, err := getDebug(flag.KubeconfigPath)
	if err != nil {
		return fmt.Errorf("get debug client: %w", err)
	}

	if podFlags.AllNamespaces {
		podFlags.Namespace = ""
	}

	if len(args) > 0 {
		fieldSelectors := strings.Split(podFlags.FieldSelector, ",")
		fieldSelectors = append(fieldSelectors, fmt.Sprintf("metadata.name=%s", args[0]))
		podFlags.FieldSelector = strings.Join(fieldSelectors, ",")
	}

	pods, err := d.Pods(ctx, podFlags.Namespace, podFlags.Label, podFlags.FieldSelector)
	if err != nil {
		return fmt.Errorf("get pods: %w", err)
	}

	if len(pods) == 0 {
		fmt.Println("found 0 pods")
	}

	pods.Print(stdOut)
	return nil
}
