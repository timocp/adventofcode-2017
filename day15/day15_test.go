package day15

import "testing"

func TestNext(t *testing.T) {
	a, b := examples()
	for i, tt := range []struct {
		expa, expb int64
	}{
		{1092455, 430625591},
		{1181022009, 1233683848},
		{245556042, 1431495498},
		{1744312007, 137874439},
		{1352636452, 285222916},
	} {
		a.next()
		b.next()
		if a.value != tt.expa {
			t.Errorf("i=%d a.next().value => %d, want %d", i, a.value, tt.expa)
		}
		if b.value != tt.expb {
			t.Errorf("i=%d b.next().value => %d, want %d", i, b.value, tt.expb)
		}
	}
}

func TestJudge(t *testing.T) {
	for _, tt := range []struct {
		aValue, bValue int64
		iterations     int64
		out            int64
	}{
		{65, 8921, 5, 1},
		{65, 8921, 40000000, 588},
	} {
		r := Judge(tt.aValue, tt.bValue, tt.iterations)
		if r != tt.out {
			t.Errorf("Judge(%d, %d, %d) => %d, want %d", tt.aValue, tt.bValue, tt.iterations, r, tt.out)
		}
	}
}

func examples() (*generator, *generator) {
	return &generator{16807, 65}, &generator{48271, 8921}
}
