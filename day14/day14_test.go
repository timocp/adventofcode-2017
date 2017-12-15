package day14

import "testing"

func TestSquaresUsed(t *testing.T) {
	in := "flqrgnkx"
	r := SquaresUsed(in)
	exp := 8108
	if r != exp {
		t.Errorf("SquaresUsed(%s) => %d, want %d", in, r, exp)
	}
}
