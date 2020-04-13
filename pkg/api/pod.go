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

func (c *Client) GetPods(namespace, labelSelector string) ([]v1.Pod, error) {

	podList, err := c.coreV1.Pods(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
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
