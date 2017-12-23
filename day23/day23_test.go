package day23

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("../input/day23.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := Part1(f)
	exp := 6241
	if r != exp {
		t.Errorf("Part1() => %d, want %d", r, exp)
	}
}
