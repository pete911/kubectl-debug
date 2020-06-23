package cmd

import (
	"fmt"
	"github.com/pete911/kubectl-debug/pkg/api"
	"github.com/pete911/kubectl-debug/pkg/debug"
	"github.com/pete911/kubectl-debug/util"
)

func getDebug(kubeconfigPath string) (*debug.Debug, error) {

	clientConfig, err := util.NewClientConfig(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("loading kubernetes client config: %w", err)
	}

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("loading kubernetes rest config: %w", err)
	}

	if podFlags.Namespace == "" {
		ns, err := util.GetNamespaceFromClientConfig(clientConfig)
		if err != nil {
			return nil, fmt.Errorf("loading namespace from client config: %w", err)
		}
		podFlags.Namespace = ns
	}

	client, err := api.NewClient(restConfig)
	if err != nil {
		return nil, fmt.Errorf("get api client: %w", err)
	}
	return debug.NewDebug(client), nil
}
