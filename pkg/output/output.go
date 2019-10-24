package output

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	Writer io.Writer = os.Stdout
)

func WriteTable(table [][]string, header bool) {

	w := new(tabwriter.Writer)
	w.Init(Writer, 0, 8, 1, '\t', tabwriter.AlignRight)
	for i, row := range table {
		if header && i == 0 {
			writeHeading(w, row)
		}
		fmt.Fprintf(w, strings.Join(row, "\t")+"\t")
	}
	w.Flush()
}

func writeHeading(w *tabwriter.Writer, row []string) {

	var underline []string
	for _, cell := range row {
		underline = append(underline, strings.Repeat("-", len(cell)))
	}
	fmt.Fprintf(w, strings.Join(row, "\t")+"\t")
	fmt.Fprintf(w, strings.Join(underline, "\t")+"\t")
}

func TruncateString(str string, num int) string {

	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}
