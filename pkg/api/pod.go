package api

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetAllPods(labelSelector string) ([]v1.Pod, error) {

	namespaces, err := c.GetNamespaces()
	if err != nil {
		return nil, fmt.Errorf("get namespaces: %w", err)
	}

	var pods []v1.Pod
	for _, namespace := range namespaces {
		p, err := c.GetPods(namespace.Name, labelSelector)
		if err != nil {
			return nil, err
		}
		pods = append(pods, p...)
	}
	return pods, nil
}

func (c *Client) GetPods(namespace, labelSelector string) ([]v1.Pod, error) {

	podList, err := c.coreV1.Pods(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return nil, err
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
