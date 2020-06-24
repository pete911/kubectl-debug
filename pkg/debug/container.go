package debug

import (
	"fmt"
	"github.com/pete911/kubectl-debug/pkg/format"
	"io"
	v1 "k8s.io/api/core/v1"
)

type Container struct {
	Name                 string
	Image                string
	Ready                bool
	Started              bool
	RestartCount         int
	State                string
	LastTerminationState string
	Logs                 []string
}

func (c Container) Print(w io.Writer) {

	fmt.Fprintf(w, "Container: %s\n", c.Name)
	fmt.Fprintf(w, "  Image:        %s\n", c.Image)
	fmt.Fprintf(w, "  Ready:        %t\n", c.Ready)
	fmt.Fprintf(w, "  Started:      %t\n", c.Started)
	fmt.Fprintf(w, "  Restart:      %d\n", c.RestartCount)
	fmt.Fprintf(w, "  State:        %s\n", c.State)
	fmt.Fprintf(w, "  Last State:   %s\n", c.LastTerminationState)
	fmt.Fprintf(w, "  Logs:")

	format.StringList(w, c.Logs, 4)
}

func toContainer(v1ContainerStatus v1.ContainerStatus, logs []string) Container {

	return Container{
		Name:                 v1ContainerStatus.Name,
		Image:                v1ContainerStatus.Image,
		Ready:                v1ContainerStatus.Ready,
		RestartCount:         int(v1ContainerStatus.RestartCount),
		Started:              toBool(v1ContainerStatus.Started),
		State:                toContainerState(v1ContainerStatus.State),
		LastTerminationState: toContainerState(v1ContainerStatus.LastTerminationState),
		Logs:                 logs,
	}
}

func toBool(b *bool) bool {

	if b == nil {
		return false
	}
	return *b
}

func toContainerState(state v1.ContainerState) string {

	if state.Running != nil {
		return "Running"
	}
	if state.Terminated != nil {
		return fmt.Sprintf("Terminated: %s", state.Terminated.Reason)
	}
	if state.Waiting != nil {
		return fmt.Sprintf("Waiting: %s", state.Waiting.Reason)
	}
	return ""
}
