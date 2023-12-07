package TUI

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"testing"
)

func TestPretty(t *testing.T) {
	t2 := table.NewWriter()
	t2.SetStyle(table.StyleDefault)
	t2.Style().Options.SeparateRows = true
	t2.SetOutputMirror(os.Stdout)

	t2.SetTitle("%20s", "Hello Title")
	t2.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t2.AppendRows([]table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	})
	t2.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	t2.AppendFooter(table.Row{"", "", "Total", 10000})
	t2.SetCaption("%s", "Hello Caption")
	t2.Render()
}
