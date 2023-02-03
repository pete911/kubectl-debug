package util

import (
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func NewClientConfig(kubeconfigPath string) (clientcmd.ClientConfig, error) {

	configBytes, err := os.ReadFile(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return clientcmd.NewClientConfigFromBytes(configBytes)
}

func GetNamespaceFromClientConfig(clientConfig clientcmd.ClientConfig) (string, error) {

	ns, _, err := clientConfig.Namespace()
	if err != nil {
		return "", err
	}
	if ns == "" {
		return "default", nil
	}
	return ns, nil
}
