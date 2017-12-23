package day22

import (
	"bytes"
	"testing"
)

var testIn = `..#
#..
...
`

func TestGrid(t *testing.T) {
	g := loadGrid(bytes.NewBufferString(testIn))
	for _, tt := range []struct {
		bursts     int
		infections int
	}{
		{7, 5},
		{70, 41},
		{10000, 5587},
	} {
		for g.bursts < tt.bursts {
			g.burst()
		}
		if g.infections != tt.infections {
			t.Errorf("After %d bursts => %d infections, want %d", g.bursts, g.infections, tt.infections)
		}
	}
}
