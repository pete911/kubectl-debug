package debug

import (
	"fmt"
	"github.com/pete911/kubectl-debug/pkg/api"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type Debug struct {
	client *api.Client
}

func NewDebug(client *api.Client) *Debug {
	return &Debug{client: client}
}

func (d *Debug) Pods(namespace, labelSelector, fieldSelector string) (Pods, error) {

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

func (d *Debug) toPods(v1Pods []v1.Pod) Pods {

	var pods []Pod
	for _, pod := range v1Pods {
		pods = append(pods, d.toPod(pod))
	}
	return pods
}

func (d *Debug) toPod(v1Pod v1.Pod) Pod {

	var containers []Container
	for _, containerStatus := range v1Pod.Status.ContainerStatuses {
		rawLogs, err := d.client.GetLogs(v1Pod.Namespace, v1Pod.Name, containerStatus.Name, 10)
		logs := strings.Split(strings.TrimSpace(string(rawLogs)), "\n")
		if err != nil {
			logs = []string{fmt.Sprintf("cannot get logs: %v", err)}
		}
		containers = append(containers, toContainer(containerStatus, logs))
	}

	v1Events, err := d.client.GetEvents(v1Pod.Namespace, v1Pod.Name)
	if err != nil {
		v1Events = []v1.Event{{Message: fmt.Sprintf("cannot get events: %v", err)}}
	}
	return toPod(v1Pod, v1Events, containers)
}
