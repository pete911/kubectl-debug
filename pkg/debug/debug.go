package debug

import "github.com/pete911/kubectl-debug/pkg/api"

type Debug struct {
	client *api.Client
}

func NewDebug(client *api.Client) *Debug {
	return &Debug{client: client}
}
