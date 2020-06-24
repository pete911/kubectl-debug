package format

import (
	"fmt"
	"io"
	"strings"
)

func StringList(w io.Writer, items []string, indent int) {

	if len(items) == 0 || (len(items) == 1 && items[0] == "") {
		fmt.Fprintf(w, "%s-\n", strings.Repeat(" ", indent))
		return
	}
	for _, item := range items {
		fmt.Fprintf(w, "%s%s\n", strings.Repeat(" ", indent), item)
	}
}

func StringMap(w io.Writer, items map[string]string, indent int) {

	if len(items) == 0 {
		fmt.Fprintf(w, "%s-\n", strings.Repeat(" ", indent))
		return
	}
	for k, v := range items {
		fmt.Fprintf(w, "%s%s: %s\n", strings.Repeat(" ", indent), k, v)
	}
}
