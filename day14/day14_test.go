package day14

import "testing"

var in = "flqrgnkx"

func TestSquaresUsed(t *testing.T) {
	r := SquaresUsed(in)
	exp := 8108
	if r != exp {
		t.Errorf("SquaresUsed(%s) => %d, want %d", in, r, exp)
	}
}

func TestRegionsPresent(t *testing.T) {
	r := RegionsPresent(in)
	exp := 1242
	if r != exp {
		t.Errorf("RegionsPresent(%s) => %d, want %d", in, r, exp)
	}
}
