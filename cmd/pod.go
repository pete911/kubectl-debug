package cmd

import (
	"fmt"
	"github.com/pete911/kubectl-debug/cmd/flag"
	"github.com/pete911/kubectl-debug/pkg/debug"
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

func runPodCmd(_ *cobra.Command, args []string) error {

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

	pods, err := d.Pods(podFlags.Namespace, podFlags.Label, podFlags.FieldSelector)
	if err != nil {
		return fmt.Errorf("get pods: %w", err)
	}

	if len(pods) == 0 {
		fmt.Println("found 0 pods")
	}

	printPods(pods)
	return nil
}

func printPods(pods []debug.Pod) {

	for _, pod := range pods {
		fmt.Printf("--- [Namespace: %s Pod: %s] ---\n", pod.Namespace, pod.Name)
		fmt.Println("Labels:")
		printMap(pod.Labels, 2)
		fmt.Println("Annotations:")
		printMap(pod.Annotations, 2)
		fmt.Println("Events:")
		printList(strings.Split(pod.Events, "\n"), 2)

		for _, container := range pod.Containers {
			fmt.Printf("Container: %s\n", container.Name)
			fmt.Printf("  Image:        %s\n", container.Image)
			fmt.Printf("  Ready:        %t\n", container.Ready)
			fmt.Printf("  Started:      %t\n", container.Started)
			fmt.Printf("  Restart:      %d\n", container.RestartCount)
			fmt.Printf("  State:        %s\n", container.State)
			fmt.Printf("  Last State:   %s\n", container.LastTerminationState)
			fmt.Println("  Logs:")

			printList(strings.Split(container.Logs, "\n"), 4)
		}
		fmt.Println()
	}
}

func printList(items []string, indent int) {

	if len(items) == 0 || (len(items) == 1 && items[0] == "") {
		fmt.Printf("%s-\n", strings.Repeat(" ", indent))
		return
	}
	for _, item := range items {
		fmt.Printf("%s%s\n", strings.Repeat(" ", indent), item)
	}
}

func printMap(items map[string]string, indent int) {

	if len(items) == 0 {
		fmt.Printf("%s-\n", strings.Repeat(" ", indent))
		return
	}
	for k, v := range items {
		fmt.Printf("%s%s: %s\n", strings.Repeat(" ", indent), k, v)
	}
}
