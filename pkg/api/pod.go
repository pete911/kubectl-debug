package api

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetAllPods(namespace string) ([]v1.Pod, error) {

	podList, err := c.coreV1.Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get all pods: %w", err)
	}
	return podList.Items, nil
}

func (c *Client) GetPods(namespace, appLabel string) ([]v1.Pod, error) {

	label := fmt.Sprintf("app=%s", appLabel)
	podList, err := c.coreV1.Pods(namespace).List(metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return nil, fmt.Errorf("get pods: %w", err)
	}
	return podList.Items, nil
}

func (c *Client) GetLogs(namespace, name, container string, tailLines int64) (string, error) {

	podLogOptions := &v1.PodLogOptions{
		Container: container,
		TailLines: &tailLines,
	}
	b, err := c.coreV1.Pods(namespace).GetLogs(name, podLogOptions).Do().Raw()
	if err != nil {
		return "", fmt.Errorf("get logs: %w", err)
	}
	return string(b), nil
}

func ContainerState(state v1.ContainerState) string {

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
