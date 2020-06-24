package debug

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type Pod struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Events      string
	Containers  []Container
}

type Container struct {
	Name                 string
	Image                string
	Ready                bool
	Started              bool
	RestartCount         int
	State                string
	LastTerminationState string
	Logs                 string
}

func (d *Debug) Pods(namespace, labelSelector, fieldSelector string) ([]Pod, error) {

	if namespace == "" {
		v1Pods, err := d.client.GetAllPods(labelSelector, fieldSelector)
		if err != nil {
			return nil, fmt.Errorf("get all pods %w", err)
		}
		return d.toPods(v1Pods), nil
	}

	v1Pods, err := d.client.GetPods(namespace, labelSelector, fieldSelector)
	if err != nil {
		return nil, fmt.Errorf("get pods %w", err)
	}
	return d.toPods(v1Pods), nil
}

func (d *Debug) toPods(v1Pods []v1.Pod) []Pod {

	var pods []Pod
	for _, pod := range v1Pods {
		pods = append(pods, d.toPod(pod))
	}
	return pods
}

func (d *Debug) toPod(v1Pod v1.Pod) Pod {

	var containers []Container
	for _, containerStatus := range v1Pod.Status.ContainerStatuses {
		logs, err := d.client.GetLogs(v1Pod.Namespace, v1Pod.Name, containerStatus.Name, 10)
		if err != nil {
			logs = fmt.Sprintf("cannot get logs: %v", err)
		}
		containers = append(containers, toContainer(containerStatus, strings.TrimSpace(logs)))
	}

	pod := Pod{
		Name:        v1Pod.Name,
		Namespace:   v1Pod.Namespace,
		Labels:      v1Pod.Labels,
		Annotations: v1Pod.Annotations,
		Containers:  containers,
	}

	v1Events, err := d.client.GetEvents(v1Pod.Namespace, v1Pod.Name)
	if err != nil {
		pod.Events = fmt.Sprintf("cannot get events: %v", err)
		return pod
	}

	pod.Events = toEvents(v1Events)
	return pod
}

func toEvents(v1Events []v1.Event) string {

	var events []string
	for _, event := range v1Events {
		events = append(events, event.Message)
	}
	return strings.Join(events, "\n")
}

func toContainer(v1ContainerStatus v1.ContainerStatus, logs string) Container {

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
