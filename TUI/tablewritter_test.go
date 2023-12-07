package TUI

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"strings"
	"testing"
)

func TestTableWriter(t *testing.T) {
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
		[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	tableTitle := ""

	tableTitle = "Some of your data"
	table.SetHeader([]string{"ID", "Name", "Status"})
	table.AppendBulk(data)
	table.Render()

	fmt.Println()
	if tableTitle != "" {
		lineLen := strings.Index(tableString.String(), "\n")
		padding := (lineLen - len(tableTitle) - 8)
		if padding < 0 {
			padding = 0
		}
		leftPadWidth := (padding / 2)
		leftPad := strings.Repeat(" ", leftPadWidth)
		rightPad := strings.Repeat(" ", padding-leftPadWidth)
		fmt.Printf("%s*** %s ***%s\n", leftPad, tableTitle, rightPad)
	}
	fmt.Println(tableString.String())
}
