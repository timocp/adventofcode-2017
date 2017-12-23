package day22

import (
	"bytes"
	"testing"
)

var testIn = `..#
#..
...
`

func TestBurst(t *testing.T) {
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

func TestEvolvedBurst(t *testing.T) {
	g := loadGrid(bytes.NewBufferString(testIn))
	for _, tt := range []struct {
		bursts     int
		infections int
	}{
		{100, 26},
		{10000000, 2511944},
	} {
		for g.bursts < tt.bursts {
			g.evolvedBurst()
		}
		if g.infections != tt.infections {
			t.Errorf("After %d evolved bursts => %d infections, want %d", g.bursts, g.infections, tt.infections)
		}
	}
}
