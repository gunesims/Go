package Format

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func printBase(s string) {
	fmt.Println(s)
}

func PrintSuccess(s string) {
	printBase(color.GreenString("[ OK ]: ") + s)
}

func PrintWarning(s string) {
	printBase(color.YellowString("[ WARNING ]: ") + s)
}

func PrintFailed(s string) {
	printBase(color.RedString("[ ERROR ]: ") + s)
}

func TestColor(t *testing.T) {
	PrintSuccess("Hello World")
	PrintWarning("Hello World")
	PrintFailed("Hello World")
}
