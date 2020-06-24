package debug

import (
	"fmt"
	"github.com/pete911/kubectl-debug/pkg/format"
	"io"
	v1 "k8s.io/api/core/v1"
	"strings"
)

type Pods []Pod

func (ps Pods) Print(w io.Writer) {

	for _, pod := range ps {
		pod.Print(w)
		fmt.Fprint(w, "\n")
	}
}

type Pod struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Events      string
	Containers  []Container
}

func (p Pod) Print(w io.Writer) {

	fmt.Fprintf(w, "--- [Namespace: %s Pod: %s] ---\n", p.Namespace, p.Name)
	fmt.Println("Labels:")
	format.StringMap(w, p.Labels, 2)
	fmt.Fprint(w, "Annotations:\n")
	format.StringMap(w, p.Annotations, 2)
	fmt.Fprint(w, "Events:\n")
	format.StringList(w, strings.Split(p.Events, "\n"), 2)

	for _, container := range p.Containers {
		container.Print(w)
	}
}

func toPod(v1Pod v1.Pod, v1Events []v1.Event, containers []Container) Pod {

	return Pod{
		Name:        v1Pod.Name,
		Namespace:   v1Pod.Namespace,
		Labels:      v1Pod.Labels,
		Annotations: v1Pod.Annotations,
		Containers:  containers,
		Events:      toEvents(v1Events),
	}
}

func toEvents(v1Events []v1.Event) string {

	var events []string
	for _, event := range v1Events {
		events = append(events, event.Message)
	}
	return strings.Join(events, "\n")
}
