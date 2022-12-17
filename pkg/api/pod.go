package api

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetAllPods(ctx context.Context, labelSelector, fieldSelector string) ([]v1.Pod, error) {

	namespaces, err := c.GetNamespaces(ctx)
	if err != nil {
		return nil, fmt.Errorf("get namespaces: %w", err)
	}

	var pods []v1.Pod
	for _, namespace := range namespaces {
		p, err := c.GetPods(ctx, namespace.Name, labelSelector, fieldSelector)
		if err != nil {
			return nil, err
		}
		pods = append(pods, p...)
	}
	return pods, nil
}

func (c *Client) GetPods(ctx context.Context, namespace, labelSelector, fieldSelector string) ([]v1.Pod, error) {

	podList, err := c.coreV1.Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: labelSelector, FieldSelector: fieldSelector})
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
}

func (c *Client) GetLogs(ctx context.Context, namespace, name, container string, tailLines int64) ([]byte, error) {

	podLogOptions := &v1.PodLogOptions{
		Container: container,
		TailLines: &tailLines,
	}
	return c.coreV1.Pods(namespace).GetLogs(name, podLogOptions).Do(ctx).Raw()
}
