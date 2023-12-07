package kingpin

import "testing"

func TestKingpin(t *testing.T) {
	var app = kingpin.New("chat", "A command-line chat application.")
}
