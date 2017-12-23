package day21

import (
	"bytes"
	"testing"
)

var testRules = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

func TestPixels(t *testing.T) {
	g := startingGrid()
	r := loadRules(bytes.NewBufferString(testRules))
	for i, tt := range []struct {
		pattern string
		pixels  int
	}{
		{"#..#/..../..../#..#", 4},
		{"##.##./#..#../....../##.##./#..#../......", 12},
	} {
		g = g.iterate(r)
		if g.String() != tt.pattern {
			t.Errorf("i=%d got %s, want %s", i, g, tt.pattern)
		}
		if g.pixels() != tt.pixels {
			t.Errorf("i=%d Pixels() => %d, want %d", i, g.pixels(), tt.pixels)
		}
	}
}
