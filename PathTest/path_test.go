package PathTest

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {

	fmt.Println(filepath.Base("./1.txt"))

	fmt.Println(filepath.Clean("C:/a/b/../c"))
}
