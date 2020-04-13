package api

import (
	"k8s.io/client-go/kubernetes"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extensionsV1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/rest"
)

type Client struct {
	networkingV1Beta1 extensionsV1.ExtensionsV1beta1Interface
	coreV1            corev1.CoreV1Interface
	batchV1           batchv1.BatchV1Interface
}

func NewClient(restConfig *rest.Config) (*Client, error) {

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		networkingV1Beta1: clientset.ExtensionsV1beta1(),
		coreV1:            clientset.CoreV1(),
		batchV1:           clientset.BatchV1(),
	}, nil
}
