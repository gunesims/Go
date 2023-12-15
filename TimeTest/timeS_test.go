package TimeTest

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ti, _ := time.Parse("Jan 02 15:04:05", "Nov 13 15:37:41")
	fmt.Println(ti)

	if ti.Year() == 0 {
		ti = ti.AddDate(time.Now().Year(), 0, 0)
	}
	fmt.Println(ti)
}
