package api

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetNamespaces(ctx context.Context) ([]v1.Namespace, error) {

	namespaceList, err := c.coreV1.Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get namespaces: %w", err)
	}
	return namespaceList.Items, nil
}
