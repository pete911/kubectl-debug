package api

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// e.g. kubectl get event --namespace abc-namespace --field-selector involvedObject.name=my-pod-zl6m6
func (c *Client) GetEvents(namespace, objectName string) ([]v1.Event, error) {

	field := fmt.Sprintf("involvedObject.name=%s", objectName)
	eventList, err := c.coreV1.Events(namespace).List(metav1.ListOptions{FieldSelector: field})
	if err != nil {
		return nil, err
	}
	return eventList.Items, nil
}
