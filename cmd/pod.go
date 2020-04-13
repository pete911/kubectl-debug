package cmd

import (
	"fmt"
	"github.com/pete911/kubectl-debug/cmd/flag"
	"github.com/pete911/kubectl-debug/pkg/api"
	"github.com/pete911/kubectl-debug/pkg/debug"
	"github.com/pete911/kubectl-debug/util"
	"github.com/spf13/cobra"
	"strings"
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
)

func init() {

	RootCmd.AddCommand(cmdPod)
	flag.InitPodFlags(cmdPod, &podFlags)
}

func runPodCmd(_ *cobra.Command, _ []string) error {

	clientConfig, err := util.NewClientConfig(flag.KubeconfigPath)
	if err != nil {
		return fmt.Errorf("loading kubernetes client config: %w", err)
	}

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("loading kubernetes rest config: %w", err)
	}

	if podFlags.Namespace == "" {
		ns, err := util.GetNamespaceFromClientConfig(clientConfig)
		if err != nil {
			return fmt.Errorf("loading namespace from client config: %w", err)
		}
		podFlags.Namespace = ns
	}

	client, err := api.NewClient(restConfig)
	if err != nil {
		return fmt.Errorf("loading api client: %w", err)
	}

	d := debug.NewDebug(client)
	pods, err := d.Pods(podFlags.Namespace, podFlags.Label)
	if err != nil {
		return fmt.Errorf("getting pods: %w", err)
	}

	if len(pods) == 0 {
		fmt.Println("found 0 pods")
	}

	printPods(pods)
	return nil
}

func printPods(pods []debug.Pod) {

	for _, pod := range pods {
		fmt.Printf("Pod: %s\n", pod.Name)
		fmt.Println("Events:")
		if pod.Events == "" {
			fmt.Printf("  -\n")
		} else {
			for _, event := range strings.Split(pod.Events, "\n") {
				fmt.Printf("  %s\n", event)
			}
		}

		for _, container := range pod.Containers {
			fmt.Printf("Container: %s\n", container.Name)
			fmt.Printf("  Image:        %s\n", container.Image)
			fmt.Printf("  Ready:        %t\n", container.Ready)
			fmt.Printf("  Started:      %t\n", container.Started)
			fmt.Printf("  Restart:      %d\n", container.RestartCount)
			fmt.Printf("  State:        %s\n", container.State)
			fmt.Printf("  Last State:   %s\n", container.LastTerminationState)
			fmt.Println("  Logs:")

			if container.Logs == "" {
				fmt.Printf("    -\n")
			} else {
				for _, log := range strings.Split(container.Logs, "\n") {
					fmt.Printf("    %s\n", log)
				}
			}
		}
		fmt.Println()
	}
}
