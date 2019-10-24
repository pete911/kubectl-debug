package debug

import (
	"fmt"
	"github.com/pete911/kubectl-debug/pkg/api"
	"github.com/pete911/kubectl-debug/pkg/output"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"strings"
	"text/tabwriter"
)

func Pod(restConfig *rest.Config) error {

	client, err := api.NewClient(restConfig)
	if err != nil {
		return fmt.Errorf("pod: %w", err)
	}

	namespaces, err := client.GetNamespaces()
	if err != nil {
		return err
	}

	var allPods []v1.Pod
	for _, namespace := range namespaces {

		pods, err := client.GetAllPods(namespace.Name)
		if err != nil {
			return fmt.Errorf("pod: %w", err)
		}
		allPods = append(allPods, pods...)
	}
	writePods(len(namespaces), allPods)
	return nil
}

func writePods(namespacesLen int, pods []v1.Pod) {

	var (
		failedPodsLen    int
		runningPodsLen   int
		succeededPodsLen int
		pendingPodsLen   int
		unknownPodsLen   int
	)

	w := new(tabwriter.Writer)
	w.Init(output.Writer, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "Namespace\tPod\tHost\tStatus\tContainers\t")
	fmt.Fprintln(w, "---------\t---\t----\t------\t----------\t")
	for _, pod := range pods {

		switch status := pod.Status.Phase; status {
		case v1.PodFailed:
			failedPodsLen++
		case v1.PodRunning:
			runningPodsLen++
		case v1.PodSucceeded:
			succeededPodsLen++
		case v1.PodPending:
			pendingPodsLen++
		default:
			unknownPodsLen++
		}

		// we are interested only in failed pods
		if pod.Status.Phase != v1.PodFailed {
			continue
		}

		var containers []string
		for _, c := range pod.Spec.Containers {
			containers = append(containers, c.Name)
		}
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%s %s %s\t%s\t",
			pod.Namespace,
			pod.Name,
			pod.Status.HostIP,
			pod.Status.Phase,
			pod.Status.Reason,
			output.TruncateString(pod.Status.Message, 50),
			strings.Join(containers, " ")),
		)
	}
	w.Flush()

	fmt.Fprintln(output.Writer)
	fmt.Fprintf(output.Writer, "Namespaces: %d Pods: %d Failed: %d Running: %d Succeeded: %d Pending: %d Unknonwn: %d\n",
		namespacesLen,
		len(pods),
		failedPodsLen,
		runningPodsLen,
		succeededPodsLen,
		pendingPodsLen,
		unknownPodsLen,
	)
}
